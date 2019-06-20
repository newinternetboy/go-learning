package main

import "fmt"

func main() {
	var age int = 25
	var expr *int
	var xiaohan_age int
	fmt.Println("hello")

	expr = &age
	xiaohan_age = *expr
	fmt.Println(expr, xiaohan_age)
}
