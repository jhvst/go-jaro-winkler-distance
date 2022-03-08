go-jaro-winkler-distance [![GoDoc](https://godoc.org/github.com/jhvst/go-jaro-winkler-distance?status.png)](https://godoc.org/github.com/jhvst/go-jaro-winkler-distance)
=====

Native [Jaro-Winkler distance](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) in Go.

Jaro-Winkler distance calculates the familiarity of two strings on range 0 to 1.

In this package the higher the score, the higher the similarity.

The package uses a sliding window for transpositions, which backtracks and looks forward for the length of the whole window. Further, the sliding window removes matches as it progresses. As a result, the values given by this package might slightly differentiate from others.

### Example

	package main

	import (
		"log"

		"github.com/jhvst/go-jaro-winkler-distance"
	)

	func main() {
        distance := Calculate("accomodate", "accommodate")
        log.Println(distance)
        // Output: 0.7509090909090909
    }
