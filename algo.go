package jwd

import (
	"math"
	"strings"
)

// According to this tool: http://www.csun.edu/english/edit_distance.php
// parameter order should have no difference in the result. Therefore,
// to avoid panicing later on, we will order the strings according to
// their length.
func order(s1, s2 string) (string, string) {
	if strings.Count(s1, "")-1 <= strings.Count(s2, "")-1 {
		return s1, s2
	} else {
		return s2, s1
	}
}

// Calculates Jaro-Winkler distance of two strings. The function lowercases and sorts the parameters
// so that that the longest string is evaluated against the shorter one.

func Calculate(s1, s2 string) float64 {

	s1, s2 = order(strings.ToLower(s1), strings.ToLower(s2))

	// m as `matching characters`
	// t as `transposition`
	// l as `the length of common prefix at the start of the string up to a maximum of 4 characters`. Not relevant in Jaro distance.
	// See more: https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
	m := 0
	t := 0
	l := 0

	window := math.Floor(float64(math.Max(float64(len(s1)), float64(len(s2)))/2) - 1)

	//debug:
	//fmt.Println("s1 param:", s1, "s2 param:", s2)
	//fmt.Println("Match window:", window, "s1:", len(s1), "s2:", len(s2))

	for i := 0; i < len(s1); i++ {
		// exact match
		if s1[i] == s2[i] {
			m++
			// Jaro-Winkler l as: `the length of common prefix at the start of the string up to a maximum of 4 characters`
			// Not relevant in Jaro distance.
			if i == l && i < 4 {
				l++
			}
		} else {
			if strings.Contains(s2, string(s1[i])) {
				// The character is considered matching if the amount of characters between the occurances in s1 and s2
				// (here `gap`) is less than match window `window`
				gap := strings.Index(s2, string(s1[i])) - strings.Index(s1, string(s1[i]))
				if gap < int(window) {
					m++
					// Somewhere in here is an error, which causes slight (max 0.04 from what I've come across, compared
					// with http://www.csun.edu/english/edit_distance.php) variations in some answers (such as slight
					// variation of Wikipedia example dicksonx and dixonsdsgese).
					// As far as I understood, transposition is only count when character exists in the string and in
					// reach of window. This loop checks whether the end of the `s2` string contains characters which
					// exist in `s1` string.
					// Editing the for loop to the following:
					//
					// 		for k := i; k < len(s1)-i; k++ {
					// 			if strings.Index(s2, string(s1[k])) < i {
					// 				fmt.Println(string(s1[k]), string(s1[i]))
					// 				t++
					// 			}
					// 		}
					//
					// Will yield same results as the libary I took example of:
					// 	https://github.com/NaturalNode/natural/blob/master/lib/natural/distance/jaro-winkler_distance.js
					// However, then the answer will be different from the ones in Wikipedia at:
					// 	https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance
					// It's a mystery to me which one is right and which one is wrong, but I'll bet on the accuracy of
					// http://www.csun.edu/english/edit_distance.php and therefore will remain the loop as following to
					// provide as accurate answer as I can.
					for k := i; k < len(s1); k++ {
						if strings.Index(s2, string(s1[k])) < i {
							t++
						}
					}
				}
			}
		}
	}

	distance := (float64(m)/float64(len(s1)) + float64(m)/float64(len(s2)) + (float64(m)-math.Floor(float64(t)/float64(2)))/float64(m)) / float64(3)
	jwd := distance + (float64(l) * float64(0.1) * (float64(1) - distance))

	//debug:
	//fmt.Println("transpositions:",tf, "m:",m)
	//fmt.Println("distance:",distance, "l:", l)
	//fmt.Println("JWD:", jwd)

	return jwd

}
