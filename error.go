package main

import "fmt"

//定义一个DivideError结构体
type DivideError struct {
	dividee int
	divider int
}

//实现`error`接口,error接口已经定义好
func (de *DivideError) Error() string {
	strFormat := `
		Cannot proceed, the divider is zero.
    		dividee: %d
    		divider: 0
	`
	return fmt.Sprintf(strFormat, de.divider)
}

//定义`int`类型除法运算函数
func Divide(varDividee int, varDivider int) (result int, errorMsg string){
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	//正常
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10= ", result)
	}
	
	//被除数非法
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)	
	}
}
