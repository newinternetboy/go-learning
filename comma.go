package main

import "fmt"

func main() {
	//单引号的类型是rune，int32的别名
	var a int = 'a'
	var b int = 10

	//双引号的类型是字符串
	var c string = "caoxiang"

	//``反单引号，创建一个原生的字符串，任何转义序列不被识别
	var d string = `\t`		
	fmt.Println(a,b,c,d)
}
