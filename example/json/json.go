package main

import "fmt"
import "encoding/json"

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("json")

	//普通类型的json格式化
	s_bool := true
	s_b_json, _ := json.Marshal(s_bool)
	fmt.Println(string(s_b_json))

	//切片json化
	sliceD := []string{"a", "b", "c"}
	sliceD_json, _ := json.Marshal(sliceD)
	fmt.Println(string(sliceD_json))

	person1 := &Person{
		Name: "caoxiang",
		Age:  25}

	person2 := Person{
		Name: "liulin",
		Age:  26}
	fmt.Println(person1)
	fmt.Println(person2)
	person1_json, _ := json.Marshal(person1)
	fmt.Println(string(person1_json))
}
