package main

import (
	"fmt"
)

func main() {
	/*
	//this will case deadclock
	message := make(chan string)
	message <- "hello"
	message <- "world"
	fmt.Println(<-message)
	fmt.Println(<-message)
	*/
	msg := make(chan string, 2)
	msg <- "hello"
	msg <- "world"
	fmt.Println(<-msg)
	fmt.Println(<-msg)
}

