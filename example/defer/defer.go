package main

import "os"
import "fmt"

func main() {
	f := createFile("./defer.txt")
	//defer 保证函数执行之后一定会执行defer后面的函数
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintf(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
