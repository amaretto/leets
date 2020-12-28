package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(mySqrt(4))
}

// use Babylonian method
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	s := float64(x / 2)
	lasts := float64(0)

	for math.Trunc(s) != math.Trunc(lasts) {
		lasts = s
		s = (s + float64(x)/s) / 2
	}
	return int(math.Trunc(s))
}

// this function use too memory
func easyMySqrt(x int) int {
	sx := math.Sqrt(float64(x))
	return int(math.Trunc(sx))
}
