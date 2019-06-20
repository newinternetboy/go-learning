package main

import (
	"fmt"
)

func loop(done chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n ", i)
	//	runtime.Gosched() // 让出cpu给其他的goroutine
	}
	done <- true
}

func main() {
	done := make(chan bool)
	for i := 0; i < 3; i++ {
		go loop(done)
	}
	for i := 0; i < 3; i++ {
		<-done
	}
	output("direct")

	go output("goroutinue")

	go func(from string) {
		for i := 0; i < 10; i++ {
			fmt.Println(from, i)
		}
	}("goroutinue anonymous")

	//enter
	fmt.Scanln()
	fmt.Println("all be done")
}

func output(from string) {
	for i := 0; i < 10 ; i++ {
		fmt.Println(from,i)
	}
}
