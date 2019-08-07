package main

import "fmt"

func main() {
	var n int = 10
	fi := fib(n)
	fmt.Println(fi)
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		fmt.Println(y)
		fmt.Println()
	}
	return y
}
