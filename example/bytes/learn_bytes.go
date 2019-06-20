package main

import (
	"bytes"
	"fmt"
)

func main() {
	s1 := "Φφϕ kKK"
	s2 := "ϕΦφ KkK"

	for _, c := range s1 {
		fmt.Printf("%-5x", c)
	}
	fmt.Println()
	for _, c := range s2 {
		fmt.Printf("%-5x", c)
	}

	s3 := []byte(" Hello World !")
	fmt.Println(bytes.Trim(s3, " !"))

	s4 := "中文"
	fmt.Println([]byte(s4))
	fmt.Println(string([]byte(s4)))
}
