package main

import "fmt"

//go 语言的初始顺序
//依赖包包级别的变量初始化>依赖包中的init函数>当前包变量初始化>当前包中init函数初始化

var b = func() int { fmt.Println("variable init in main"); return 2 }()

func init() {
    fmt.Println("main init")
}

func main() {

}
