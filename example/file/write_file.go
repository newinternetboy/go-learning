package main

import "fmt"

import (
	"bufio"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//这个ioutil写入
	str_bytes := []byte("hello world!")
	err := ioutil.WriteFile("/tmp/test1.txt", str_bytes, 0644)
	check(err)

	//读取数据
	data, err := ioutil.ReadFile("/tmp/test1.txt")
	check(err)
	fmt.Println(data)

	//io包创建文件写入文件
	f, err := os.Create("/tmp/test2.txt")
	check(err)
	//使用defer在函数执行完毕后关闭句柄
	defer f.Close()

	//写入[]byte
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//直接写入字符串
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	//将所有缓存刷出，写入文件
	f.Sync()

	//使用bufio包缓冲写入数据
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	//刷出缓冲，写入磁盘
	w.Flush()
}
