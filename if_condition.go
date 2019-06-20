package main

import (
	"fmt"
	"math"
)

func main(){
	var a int = 10
	
	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 的值为:%d\n",a)
	}
	
	var grade string = "A"
	//switch的没有必要是常量，值也不必要是整形；
	//switch没有case项等同于switch true
	switch {
		case grade == "A" :
			fmt.Printf("优秀\n")
		case grade == "B" :
			fmt.Printf("良好\n")
		default:
			fmt.Print("差\n")
	}
	fmt.Println(pow(3, 2, 10),pow(3, 3, 20))
}


func pow(x, n, lim float64) float64{
	//在条件执行之前，可以执行表达式进行运算,v可以在if结构体中使用，但是不可以在结构体外使用。
	if v := math.Pow(x, n); v < lim {
		return v
	}
	//can`t use v here,though
	return lim
}
