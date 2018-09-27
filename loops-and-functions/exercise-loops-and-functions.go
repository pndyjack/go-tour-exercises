package main

import (
	"fmt"
	"math"
)

// Sqrt unexported
func Sqrt(num float64) float64 {
	root := 10.0
	for math.Abs(getRoot(num, root)-root) > .001 {
		root = getRoot(num, root)
	}
	return root
}

func getRoot(num, root float64) float64 {
	return root - ((root*root - num) / (2 * root))
}

func main() {
	for index := 2.0; index < 11; index++ {
		fmt.Println(Sqrt(index) * Sqrt(index))
	}
}
