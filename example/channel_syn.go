package main

import (
	"fmt"
	"time"
)

func worker(done chan bool){
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
	fmt.Println("main is done")
}
func main() {
	done := make(chan bool, 1)
	go worker(done)
	fmt.Println(<-done)
	chan1 := make(chan bool, 1)
	worker1(chan1);
	value := <-chan1
	fmt.Println(value)
}


func worker1(chan1 chan bool){
	fmt.Print("chan1 is working")
	chan1 <- true
	fmt.Print("chan1 has got message")
}
