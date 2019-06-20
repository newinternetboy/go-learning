package main

import "fmt"

//slice 不存储任何数据，只是引用原数据的一部分。slice对应的修改也会作用到原数据
//slice 的声明和数组的声明很类似，只是没有相应的长度
func main() {
	// []表示切边， []type => []int 表示切片中元素为int类型
	// 3,表示切片初始长度
	// 5,表示切片的容量
	var slice []int
	if slice == nil {
		fmt.Println("slice don`t be initialized and value is nil")
	}

	slice1 := make([]int, 3, 5)
	getMetaInfoOfSlice(slice1)
	var i int
	for i = 0; i < 10; i++ {
		slice1 = append(slice1, i)
	}
	fmt.Println(slice1)

	var slice2 []int
	//创建一个切片slice2容量是slice1的两倍
	slice2 = make([]int, len(slice1), cap(slice1)*2)
	getMetaInfoOfSlice(slice2)
	//赋值slice1的元素到slice2中
	copy(slice2, slice1)
	getMetaInfoOfSlice(slice2)

	//++++++++++++++++++++//
	names := [3]string{
		"CAO",
		"LIU",
		"ZHANG",
	}

	names1 := names[1:2]
	names1[0] = "LIU XIAO HAN"
	fmt.Println(names[1])

	//多维slice，slice的长度是多变的
	twoD := make([][]int, 3)
	for k := 0; k < 3; k++ {
		innerLen := k + 1
		twoD[k] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[k][j] = k + j
		}
	}
	fmt.Println("2d:", twoD)
}

func getMetaInfoOfSlice(x []int) {
	fmt.Printf("切片长度len=%d 容量cap=%d slice=%v\n", len(x), cap(x), x)
}
