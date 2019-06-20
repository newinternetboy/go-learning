package main

import "encoding/json"
import "fmt"

//定义小写的成员结构体
type Person struct {
	name string
	age  int
}

//结构体实现对应的MarshalJSON()接口
func (this Person) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name": this.name,
		"age":  this.age})
}

func main() {
	person := Person{
		name: "caoxiang",
		age:  25}

	p_struct_json, _ := json.Marshal(person)

	fmt.Println(string(p_struct_json))
}
