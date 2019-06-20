package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second * 1)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
