package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for k, v := range os.Args[1:] {
		fmt.Printf("%d %s\n", k, v)
	}
	//耗时
	fmt.Printf("%.2f \n", time.Since(start).Seconds())
}
