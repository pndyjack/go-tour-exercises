package main

import (
	"io"
	"os"
	"strings"
	"unicode"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(cipherText []byte) (int, error) {
	length, err := r13.r.Read(cipherText)
	if err == io.EOF {
		return 0, err
	}
	for index := range cipherText[:length] {
		if unicode.IsLetter(rune(cipherText[index])) {
			switch unicode.IsLower(rune(cipherText[index])) {
			case true:
				cipherText[index] -= 97
				cipherText[index] += 13
				cipherText[index] %= 26
				cipherText[index] += 97
			case false:
				cipherText[index] -= 65
				cipherText[index] += 13
				cipherText[index] %= 26
				cipherText[index] += 65
			}
		}
	}
	return length, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
