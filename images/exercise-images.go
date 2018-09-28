package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

// Image unexported
type Image struct {
	w   int
	h   int
	img [][]uint8
}

// ColorModel unexported
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds unexported
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At unexported
func (i Image) At(x, y int) color.Color {
	v := i.img[x][y]
	return color.RGBA{
		v,
		v,
		255,
		255,
	}
}

func main() {
	m := Image{
		w:   255,
		h:   255,
		img: Pic(255, 255),
	}
	pic.ShowImage(m)
}

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
