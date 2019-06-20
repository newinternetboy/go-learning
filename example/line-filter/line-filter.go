package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//使用bufio.NewScanner()可以将无缓冲的`os.Stdin`包裹，以便使用`Scan`方法
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	//错误检查
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
