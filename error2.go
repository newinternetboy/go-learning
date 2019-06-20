package main

import (
	"fmt"
	"errors"
)

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error) {
	if arg == 0 {
		return arg, errors.New("number innagative")
	}
	
	return arg, nil
}

func f2(arg int) (int, error){
	if arg == 0 {
		return arg, &argError{arg, "number is zero"}
	}
	
	return arg, nil
}

func main() {
	i := 0
	if arg,error:= f1(i); error != nil {
		fmt.Println("f1 failed:", error)
	} else {
		fmt.Println("f1 worked:",arg )
	}
	
	if r, e := f2(i); e != nil {
		fmt.Println("f2 failed:", e)
	}else {
		fmt.Println("f2 worked:", r)
	}

	er := errors.New("this is error new")
	fmt.Println(er)
}
