package main

import "fmt"

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//创建接受信号的缓冲隧道
	signs := make(chan os.Signal, 1)
	//监测程序执行结束的隧道
	done := make(chan bool)

	//将程序接受到的信号绑定到接受信号的通道上
	signal.Notify(signs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sign := <-signs
		fmt.Println(sign)
		//通知主线程协程执行完毕
		done <- true
	}()

	fmt.Println("wating signal...")
	<-done
	fmt.Println("accept signal...")
}
