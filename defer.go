package main

import "fmt"

//函数执行的顺序,先执行main函数，main函数执行完毕，执行defer关键字修饰的函数
func main() {
	defer world()

	fmt.Println("Hello ")

	//stacking defer(栈defer)
	//多个defer，后进先出
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("main function has done!")
}

func world() {
	fmt.Println("Worlds!")
}
