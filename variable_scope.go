package main

import "fmt"

func main() {
	//if语句创建的变量对后面的分支if语句是可见的
	if x := 1; x == 0 {
		fmt.Println(x)
	} else if y := 1; x == y {
		fmt.Println("x variable is visible")
	}
}
