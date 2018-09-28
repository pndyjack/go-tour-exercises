package main

import (
	"fmt"
	"math"
)

const delta = 0.001

const initialRoot = 10.0

// ErrNegativeSqrt unexported
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt unexported
func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNegativeSqrt(num)
	}
	root := initialRoot
	for math.Abs(getRoot(num, root)-root) > delta {
		root = getRoot(num, root)
	}
	return root, nil
}

func getRoot(num, root float64) float64 {
	return root - ((root*root - num) / (2 * root))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
