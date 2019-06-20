package main

import "os"

func main() {
	//	panic("this is a imnormal problem!")

	_, error := os.Create("/tmp/test.txt")
	if error != nil {
		panic("you need to fix this error in proper way")
	}
}
