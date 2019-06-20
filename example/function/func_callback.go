package main

import "fmt"

type Callback func(x, y int) int

func main() {
	x, y := 1, 2
	fmt.Println(test(x, y, add))
}

func add(x, y int) int {
	return x + y
}

//定义回调函数
func test(x, y int, callback Callback) int {
	return callback(x, y)
}
