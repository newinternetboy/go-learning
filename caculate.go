package main

import "fmt"

func main() {
	//除法，浮点数相除则得浮点数，整数相除舍去浮点数取整
	x, y := 1.0, 2.0
	fmt.Println(x / y)
	x1, y1 := 1, 2
	fmt.Println(x1 / y1)
	//算数表达式进行运算如果出现溢出，那么会舍弃高位
	var u uint8 = 255
	fmt.Println(u + 1)
}
