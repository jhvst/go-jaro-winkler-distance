package jwd

import (
	"bytes"
	"math"
	"strings"
	"unicode/utf8"
)

func sort(s1, s2 string) (shorter, longer string) {
	if utf8.RuneCountInString(s1) < utf8.RuneCountInString(s2) {
		return s1, s2
	}
	return s2, s1
}

// Calculate calculates Jaro-Winkler distance of two strings.
// The function lowercases its parameters.
func Calculate(s1, s2 string) float64 {

	s1, s2 = sort(strings.ToLower(s1), strings.ToLower(s2))

	// A sliding window for match search.
	window := uint(math.Floor(
		math.Max(
			float64(utf8.RuneCountInString(s1)),
			float64(utf8.RuneCountInString(s2)),
		)/2,
	) - 1)

	runes1 := bytes.Runes([]byte(s1))
	runes2 := bytes.Runes([]byte(s2))

	var m uint = 0
	var matches []bool
	for i := 0; i < len(runes1); i++ {
		match := false
		if runes1[i] == runes2[i] {
			m++
			match = true
		}
		matches = append(matches, match)
	}

	if m == 0 {
		return 0.0
	}

	var t float64 = 0
	slider := runes2[0:window]
	for i := 0; i < len(runes1); i++ {

		if matches[i] {
			continue
		}

		idx := strings.Index(string(slider), string(runes1[i]))
		if idx != -1 && !matches[idx] {
			t += 0.5
			matches[idx] = true
		}

		start := uint(math.Max(
			0,
			float64(i-int(window)),
		))
		end := uint(math.Min(
			float64(i+int(window)),
			float64(len(runes1)),
		))

		slider_new := runes2[int(start):int(end)]
		if len(slider_new) >= int(window) {
			slider = slider_new
		}
	}

	var term1, term2, term3 float64
	term1 = float64(m) / float64(len(runes1))
	term2 = float64(m) / float64(len(runes2))
	term3 = (float64(uint(float64(m) - t))) / float64(m)

	var simj float64
	simj = (term1 + term2 + term3) / 3

	var p float64 = 0.1
	var l uint = 0
	var common_prefix uint = uint(math.Min(4.0, float64(len(s1))))
	for i := 0; i < len(s1[0:common_prefix]); i++ {
		if s1[0:common_prefix][i] == s2[0:common_prefix][i] {
			l++
		}
	}

	return simj + float64(l)*p*(1-simj)
}
