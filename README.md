go-jaro-winkler-distance [![GoDoc](https://godoc.org/github.com/9uuso/go-jaro-winkler-distance?status.png)](https://godoc.org/github.com/9uuso/go-jaro-winkler-distance)
=====

Native [Jaro-Winkler distance](https://en.wikipedia.org/wiki/Jaro%E2%80%93Winkler_distance) in Go.

Jaro-Winkler distance calculates the familiriaty of two strings on range 0 to 1.

For example comparing words `DIXON` and `DICKSONX` gives you score of `0.8133333333333332`, whilst comparing words `sound` and `ääni` will yield `0.438`. That being said, this package also supports unicode characters.

### Example

	package main

	import (
		"fmt"

		"github.com/9uuso/go-jaro-winkler-distance"
	)

	func main() {
		// See more example strings at http://www.amstat.org/sections/srms/Proceedings/papers/1990_056.pdf
		fmt.Println(jwd.Calculate("DIXON", "DICKSONX"))
		// output: 0.8133333333333332
	}
