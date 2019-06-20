package main

import "fmt"

func main(){
	//Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素
	//for 循环中使用range作用切片，返回索引=>元素值
	//弃用索引_,如果只需要索引可以只写一个参数
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums{
		sum += num
	}
	fmt.Println(sum)
	
	//需要索引
	for i, num := range nums{
		if num == 3 {
			fmt.Println("index", i)
		}
	}

	//range用来枚举Unicode字符串，索引=>字符对应的unicode值本身
	for i,s := range "iloveyou" {
		fmt.Println(i,s)
	}
	
	//遍历map
	//省略了value
	maps := map[string]string{"JYF":"陀思妥耶夫斯基","MSLDZL":"HD"}
	for k := range maps {
		fmt.Println(k)
	}
}
