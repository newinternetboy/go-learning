package main

import "fmt"

func main() {

	//map不像其他数据类型(没有初始化系统自动初始化),最好使用make进行初始化map
	//map_name := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	map_name := make(map[string]string)
	map_name["France"] = "Praris"
	map_name["china"] = "Beijing"
	map_name[""] = ""
	fmt.Println(map_name)
	f
	_, value := map_name["china"]
	fmt.Println(value)
	_, value1 := map_name["china2"]
	fmt.Println(value1)
	_, key_exist := map_name[""]
	fmt.Println(key_exist)
	//遍历返回键=>值
	for country, city := range map_name {
		//		fmt.Println(country,"首都是", map_name[country])
		fmt.Println(country, city)
	}
}
