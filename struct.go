package main

import "fmt"

type Books struct {
	title string
	author string
	book_id int
}

type Person struct {
	name string
	age  int
}
type rect struct {
        width, height int
}

//接受的是指针
func (r *rect) area() int {
        return r.width * r.height
}
//接受的是struct实参
func (r rect) zc() int {
        return 2*r.width + 2*r.height
}
func main(){
	var books Books
	books = Books{"Golang","somebody",11}
	fmt.Println(books)
	//结构体成员的访问，结构体对象.成员名
	fmt.Println(books.title)
	
	//结构体指针
	var p_books *Books
	p_books = &books
	//结构体指针访问成员
	fmt.Println(p_books.title)
	
	person := Person{name:"caoxiang", age:26}
	fmt.Println(person)
	//数据读取
	fmt.Println(person.name)
	//修改值
	person.name = "hxh"
	fmt.Println(person)
	//指针操作
	personp := &person
	fmt.Println(personp.name,personp.age)

	rect := rect{width:10, height:10}
	fmt.Println("area:", rect.area())
	fmt.Println("zc:", rect.zc())
	rectp := &rect
	fmt.Println("area:", rectp.area())
	fmt.Println("zc:", rectp.zc())
}

