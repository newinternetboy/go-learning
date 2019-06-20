package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("Hello,Go.This is %d\n", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("Hello,Go.This is %d\n", i)
		}
	}()
	//当计数器大于0的话，main()函数会阻塞
	wg.Wait()
}
