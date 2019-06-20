package main

import "fmt"

type argError struct{
	arg int
	subscribe string
}

//参数结构体实现error接口
func(e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.arg, e.subscribe)
}

func main() {
	f := 0
	if f == 0 {
		//returns it as an error created by errors.New.
		e := fmt.Errorf("number can`t be %d",f)
		fmt.Println(e)
	}
	if _, e := f1(0); e != nil {
		fmt.Println("f1 failed:", e)
	}
}

func f1(arg int)(int, error){
	if arg == 0 {
		return arg,&argError{arg, "can`t be zero"}
	}
	return arg, nil
}
