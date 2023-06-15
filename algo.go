package jwd

import (
	"bytes"
	"math"
	"unicode/utf8"
)

func reduce[T, M any](s []T, f func(M, T) M, init M) M {
	acc := init
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func mkRange[T any](slice []T, min, max, lo, hi int) []T {
	s := uint(math.Max(float64(min), float64(lo)))
	e := uint(math.Min(float64(hi), float64(max)))
	return slice[s:e]
}

func trans(t int, k int, v rune, p1 int, runes2 []rune, matches []bool) {
	slider := mkRange(runes2,
		0, p1,
		k-t, k+t+1,
	)
	for kk, vv := range slider {
		if matches[k-t+kk] {
			continue
		}
		matches[k-t+kk] = vv == v
	}
}

func non_trans(shorter, longer string, window uint) ([]bool, []bool) {
	matches := make([]bool, utf8.RuneCountInString(shorter))
	transpositions := make([]bool, utf8.RuneCountInString(shorter))

	slider := mkRange(bytes.Runes([]byte(shorter)),
		0, len(shorter),
		0, len(shorter),
	)
	for k, v := range slider {

		t := reduce(transpositions, func(acc int, current bool) int {
			if current {
				acc += 1
			}
			return acc
		}, 0)
		if t > 0 {
			trans(t, k, v, utf8.RuneCountInString(shorter), bytes.Runes([]byte(longer)), matches)
		}

		if matches[k] {
			continue
		}
		matches[k] = v == bytes.Runes([]byte(longer))[k]

		idx := bytes.IndexRune(mkRange([]byte(longer),
			0, utf8.RuneCountInString(shorter),
			k-int(window), k+int(window),
		), v)
		if idx != -1 && !matches[idx] {
			transpositions[k] = true
		}
	}
	return matches, transpositions
}

// Calculate calculates Jaro-Winkler distance of two strings.
// The function lowercases its parameters.
func Calculate(shorter, longer string) float64 {

	if utf8.RuneCountInString(shorter) > utf8.RuneCountInString(longer) {
		return Calculate(longer, shorter)
	}

	// A sliding window for match search.
	window := uint(math.Floor(float64(utf8.RuneCountInString(longer))/2) - 1)
	matches, transpositions := non_trans(shorter, longer, window)

	m := reduce(matches, func(acc int, current bool) int {
		if current {
			acc += 1
		}
		return acc
	}, 0)
	t := reduce(transpositions, func(acc int, current bool) int {
		if current {
			acc += 1
		}
		return acc
	}, 0)

	var term1, term2, term3 float64
	term1 = float64(m) / float64(utf8.RuneCountInString(shorter))
	term2 = float64(m) / float64(utf8.RuneCountInString(longer))
	term3 = (float64(uint(float64(m) - float64(t)/2))) / float64(m)

	var simj float64
	simj = (term1 + term2 + term3) / 3

	var p float64 = 0.1
	var l uint = 0
	var common_prefix uint = uint(math.Min(4.0, float64(len(shorter))))
	for i := range shorter[0:common_prefix] {
		if shorter[0:common_prefix][i] == longer[0:common_prefix][i] {
			l++
		}
	}

	return simj + float64(l)*p*(1-simj)
}
