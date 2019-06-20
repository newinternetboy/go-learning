package main

import (
	"fmt"
	"math"
)

//定义一个接口
type Phone interface {
	call()
}

//定义一个结构体
type NokiaPhone struct {
}

//结构体NokiaPhone 实现call()方法
func (nokia NokiaPhone) call() {
	fmt.Printf("I`m nokia,i can call")
}

// 算数接口定义
type geometry interface {
	area() float64
	perim() float64
}

// 定义结构体矩形
type rect struct {
	width, height float64
}

//定义结构体圆
type circle struct {
	radius float64
}

//矩形结构体实现geometry接口
func (r rect) area() float64 {
	return r.width * r.height
}

 func (r rect) perim() float64 {
	return 2*(r.width + r.height)
}

//圆结构体实现geometry接口
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return math.Pi * 2 * c.radius
}
func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	rect := rect{width:10,height:10}
	circle := circle{radius:2}
	
	fmt.Println("rect area:", rect.area())
	fmt.Println("rect perim:", rect.perim())

	fmt.Println("circle area", circle.area())
	fmt.Println("circle perim", circle.perim())
}
