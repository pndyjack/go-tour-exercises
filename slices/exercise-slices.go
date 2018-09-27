package main

import (
	"golang.org/x/tour/pic"
)

// Pic unexported
func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := range picture {
		picture[i] = make([]uint8, dx)
	}
	for i, row := range picture {
		for j := range row {
			row[j] = uint8((i ^ j))
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
