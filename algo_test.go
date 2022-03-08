package jwd

import (
	"log"
)

func ExampleCalculate() {
	distance := Calculate("accomodate", "accommodate")
	log.Println(distance)
	// Output: 0.7509090909090909
}
