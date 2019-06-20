package main

import "fmt"

func main() {
	i := 1
	in, sum := sum_self(i)
	fmt.Println(in, sum)
}

func sum_self(x int) (a, s int) {
	s = x + x
	a = x
	return
}
