package main

import (
	"fmt"
	"math"
)

func main() {
	hyp := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hyp(4,5))
	fmt.Println(doMe(hyp))
}

func doMe(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}