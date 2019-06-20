package main

import "fmt"

type StringError struct{
	str string
	msg string
}

//用该结构体实现error接口
func (e *StringError) Error() string {
	return fmt.Sprintf("%s - %s", e.str, e.msg)
}
func main() {
	name := "a"
	if name != "b" {
		e := &StringError{name, "invalid string"}
		fmt.Println(e)
	}
}
