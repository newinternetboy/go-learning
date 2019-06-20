package main

import "fmt"

func main() {
	var arr [10] int
	var i, j int
	/* 初始化数组 */
	for i=0; i < 10; i++ {
		arr[i] = i + 100
	}
	
	/* 取出数组的值 */
	for j=0; j < 10; j++ {
		fmt.Printf("索引为%d的数值为%d\n",j,arr[j])
	}

	/* 多维数组的值得获取 */
	var x int = 3
	var y int = 3
	var many_arr [3][3]int
	many_arr = many_array()
	for x=0; x < 3; x++{
		for y=0; y < 3; y++{
				fmt.Printf("%d*%d = %d\n",x, y, many_arr[x][y])
		}
	}
}

//多维数组的初始化 var array_name [x][y][z] variable_type

func many_array() [3][3]int {
	// 数组定义的时候，不能直接使用变量
	var array [3][3] int
	var i, j int
	for i=0; i < 3; i++ {
		for j=0; j < 3; j++ {
				array[i][j] = i * j 
		}
	}
	return array
}

