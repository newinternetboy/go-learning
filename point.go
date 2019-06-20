package main

import "fmt"

func main(){
	var ptr *int
	if ptr == nil {
		fmt.Printf("空指针\n")
	}else{
		fmt.Printf("不是空指针\n")
	}
	
	var arr = [3]int{10,20,30}
	//数组指针,其实和正常的数组声明也差不了多少呢
	var arr_point [3] *int
	var i int
	for i=0; i < 3 ; i++ {
		arr_point[i] = &arr[i]
		fmt.Printf("数组下标为%d的地址为%d\n",i,&arr[i])
	}

	//指针的指针
	var point_ptr **int
	var a int = 100
	ptr = &a
	point_ptr = &ptr
	fmt.Printf("ptr的指针变量的值%d\n",ptr)
	fmt.Printf("ptr的指针变量的指针%d\n",point_ptr)

	//简单数据类型的指针直接指向数值本身
	age := 26
	fmt.Printf("%d",age)

	//指针获取值，并通过*引用操作符设置变量的值
	var b int
	var p *int
	b = 22
	//取地址符&
	p = &b
	fmt.Printf("地址:%d\n",p)
	//*引用符，*p指向b的地址
	*p = 23
	fmt.Printf("b change:%d\n",b)


	/* 指针 &取地址符 *地址引用 */
	v := 0
	fmt.Println(v)
	changevalue(v)
	fmt.Println(v)
	changevalueByPoint(&v)
	fmt.Println(v)
}

func changevalue(v int) {
	v = 1
}

func changevalueByPoint(vp *int) {
	*vp = 1
}
