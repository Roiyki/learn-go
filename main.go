package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")

	x := math.Exp(3)
	fmt.Printf("number %g is:\n", x)

	y := inc(conv(x))
	fmt.Printf("number %d is:\n", y)

	y = inc(y)
	fmt.Printf("number %d is:\n", y)
}

func conv(x float64) int {
	return int(x)
}

func inc(x int) int {
	return x + 1
}
