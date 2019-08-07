package main

import "fmt"

func main() {
	//每次调用函数都会用一个新的地址存储v
	fmt.Println(f())
	fmt.Println(f())
}

func f() *int {
	v := 1
	return &v
}
