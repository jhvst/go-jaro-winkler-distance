package jwd

import (
	"log"
)

func ExampleCalculate() {
	distance := Calculate("accomodate", "accommodate")
	log.Println(distance)
	// Output: 0.8387012987012987
}

func ExampleCalculate4() {
	distance := Calculate("FAREMVIEL", "FARMVILLE")
	log.Println(distance)
	// Output: 0.6824074074074075
}

func ExampleCalculate3() {
	distance := Calculate("accommodate", "accomodate")
	log.Println(distance)
	// Output: 0.7509090909090909
}

func ExampleCalculate2() {
	distance := Calculate("asdfadsfT", "Academic Free License v1.2")
	log.Println(distance)
	// Output: 0.04985754985754986
}
