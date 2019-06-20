package main

import "fmt"
import "time"

func main() {
	c := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		c <- true
	}()
	time.Sleep(time.Second * 2)
	//上述的管道发送了一个true，但是main没有接收，此时main函数会发生阻塞
	a := make(chan bool)
	a <- true
}
