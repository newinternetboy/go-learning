package main

import "fmt"

func main() {
	//数组
	arr := []int{1, 2, 3}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	//字典
	map_string := map[string]string{
		"a": "aa",
		"b": "bb",
		"c": "cc"}
	for mk, mv := range map_string {
		fmt.Println(mk, mv)
	}

	//字符串
	range_string := "hello world"
	for _, sv := range range_string {
		fmt.Println(string(sv))
	}
}
