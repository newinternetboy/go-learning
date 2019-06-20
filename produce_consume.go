package main

import (
	"fmt"
	"time"
)

func main() {
	bread := make(chan int, 3)
	for i := 1; i <= 2; i++ {
		go produce(bread)
	}
	for i := 1; i <= 5; i++ {
		go consume(bread)
	}
	//主线程在该时间内2位面包师生产面包，5个顾客在取。
	time.Sleep(1e10)
}

func produce(ch chan<- int) {
	for {
		ch <- 1
		fmt.Println("produce bread")
		time.Sleep(100 * time.Millisecond)
	}
}

func consume(ch <-chan int) {
	for {
		<-ch
		fmt.Println("take bread")
		time.Sleep(200 * time.Millisecond)
	}
}
