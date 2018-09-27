package main

import "fmt"

func fibonacci() func(int) int {
	return func(num int) int {
		f0, f1 := 0, 1
		switch {
		case num == 0:
			return f0
		case num == 1:
			return f1
		}
		var res int
		for index := 2; index <= num; index++ {
			res = f0 + f1
			f0 = f1
			f1 = res
		}
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
