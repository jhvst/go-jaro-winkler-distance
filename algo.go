package jwd

import (
	"bytes"
	"math"
	"strings"
	"unicode/utf8"
)

var Weight float64 = 0.1
var CommonPrefixLimitter int = 4

// To avoid panicing later on, order strings according to
// their unicode length.
func sort(s1, s2 string) (shorter string, longer string) {
	if utf8.RuneCountInString(s1) < utf8.RuneCountInString(s2) {
		return s1, s2
	}
	return s2, s1
}

func window(s2 string) float64 {
	return math.Floor(float64(utf8.RuneCountInString(s2)/2) - 1)
}

func score(m, t, runes1len, runes2len float64) float64 {
	return (m/runes1len + m/runes2len + (m-math.Floor(t/float64(2)))/m) / float64(3)
}

// Calculate calculates Jaro-Winkler distance of two strings.
// The function lowercases its parameters.
func Calculate(s1, s2 string) float64 {

	// Avoid returning NaN
	if utf8.RuneCountInString(s1) == 0 || utf8.RuneCountInString(s2) == 0 {
		return 0
	}

	s1, s2 = sort(strings.ToLower(s1), strings.ToLower(s2))

	// m as `matching characters`
	// t as `transposition`
	// l as `the length of common prefix at the start of the string up to a maximum of 4 characters`.
	// See more: https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
	m := float64(0)
	t := float64(0)
	l := 0

	window := window(s2)

	runes1 := bytes.Runes([]byte(s1))
	runes2 := bytes.Runes([]byte(s2))

	for i := 0; i < len(runes1); i++ {

		// Exact match
		if runes1[i] == runes2[i] {
			m++
			// Common prefix limitter
			if i == l && i <= CommonPrefixLimitter {
				l++
			}
		} else {
			if strings.Contains(s2, string(runes1[i])) {
				// The character is also considered matching if the amount of characters between the occurances in s1 and s2
				// is less than match window
				gap := float64(strings.Index(s2, string(runes1[i])) - strings.Index(s1, string(runes1[i])))
				if gap <= window {
					m++
					// Check if transposition is in reach of window
					for k := i; k < len(runes1); k++ {
						if strings.Index(s2, string(runes1[k])) <= i {
							t++
						}
					}
				}
			}
		}
	}

	score := score(m, t, float64(len(runes1)), float64(len(runes2)))
	distance := score + (float64(l) * Weight * (float64(1) - score))

	//debug:
	//fmt.Println("- score:", score)
	//fmt.Println("- transpositions:", t)
	//fmt.Println("- matches:", m)
	//fmt.Println("- window:", window)
	//fmt.Println("- l:", l)
	//fmt.Println("- distance:", distance)

	return distance
}
