package main

import "fmt"

func main() {
	slice_data := make([]string, 3)
	sl, sc := getSliceLenCap(slice_data)
	fmt.Println(sl, sc)

	//常规操作slice，插入数据
	slice_data[0] = "a"
	slice_data[1] = "b"
	slice_data[2] = "c"
	fmt.Println(slice_data)

	//使用append()方法新增元素,返回新的切片
	slice_data1 := append(slice_data, "d")
	fmt.Println(slice_data1)
	fmt.Println(slice_data)

	//使用copys(d, g)将g切片拷贝到d切片中
	//1 d的长度小于g  只会拷贝d的长度个数的元素到d中
	slice_data2 := make([]string, 3)
	copy(slice_data2, slice_data1)
	fmt.Println(getSliceLenCap(slice_data2))

	//2 d的长度大于等于g
	slice_data3 := make([]string, len(slice_data2))
	copy(slice_data3, slice_data2)
	fmt.Println(slice_data3)
	fmt.Println(getSliceLenCap(slice_data3))

	//slice[low,high],[low,high)
	fmt.Println(slice_data3[:])
	fmt.Println(slice_data3[:1])
	fmt.Println(slice_data3[2:])
	fmt.Println(slice_data3[:len(slice_data3)])

	//多维切片
	twoD := make([][]int, 3)
	fmt.Println(len(twoD))
	for i := 0; i < len(twoD); i++ {
		inner_len := i + 1
		twoD[i] = make([]int, inner_len)
		for j := 0; j < inner_len; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}

func getSliceLenCap(slice []string) (l, c int) {
	return len(slice), cap(slice)
}
