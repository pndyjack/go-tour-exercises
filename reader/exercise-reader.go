package main

import (
	"golang.org/x/tour/reader"
)

// MyReader unexported
type MyReader struct{}

func (mr MyReader) Read(b []byte) (int, error) {
	for i := range b[:len(b)] {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
