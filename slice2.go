package main

import (
	"fmt"
	"strings"
)

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	//slice 可以包含任何数据类型，包括slice本省 [][]string定义了二维切片
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//change value
	board[0][0] = "X"
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	//slice 可以通过内置函数append添加元素，当返回的数组比原定义数组的长度还大，会申请新的内存，数组指针指向新的内存地址
	array1 := [3]int{}
	a_slice1 := array1[:]
	printSlice("a_slice1", a_slice1)
	a_slice1 = append(a_slice1, 1, 2, 3)
	printSlice("a_slice1", a_slice1)
	a_slice1 = append(a_slice1, 4)
	printSlice("a_slice1", a_slice1)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v address=%d\n", s, len(x), cap(x), x, &x)
}
