package main

import "fmt"

func main() {
	//在go语言中没有初始化的变量，系统会根据对应的类型进行初始化
	numbers := [6]string{"a", "b", "c"}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x的值 = %s\n", i, x)
	}
}
