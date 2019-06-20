//dup2 打印输入中多次出现的行的个数和文本
//它从stdin 或 指定的文件列表读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		//从文件列表中读取
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}
			//countLines(f, counts)
			input := bufio.NewScanner(f)
			for input.Scan() {
				counts[input.Text()]++
				if counts[input.Text()] > 1 {
					fmt.Println(arg)
					break
				}
			}
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	i := 1
	for input.Scan() {
		counts[input.Text()]++
		i++ 
		if i > 6 {
			break
		}
	}	
}
