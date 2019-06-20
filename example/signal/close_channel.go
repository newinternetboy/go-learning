package main

import "fmt"

func main() {
	//创建缓冲通道
	data := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			v, err := <-data
			if err {
				fmt.Println(v)
			} else {
				fmt.Println("data channel is empty or closed!")
				done <- true
			}
		}
	}()

	//向通道中写入数据
	for i := 0; i < 3; i++ {
		data <- i
	}

	//关闭通道
	close(data)

	<-done
}
