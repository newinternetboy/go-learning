package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	r, _ := regexp.Compile("p([a-z]+)ch")
	
	//该正则可以匹配到，能匹配到返回true，反之返回false
	fmt.Println(r.MatchString("peach"))
	
	//r.FindString() 返回匹配到的字符串
	fmt.Println(r.FindString("peach punch"))

	//r.FindStringIndex()
	fmt.Println(r.FindStringIndex("peach punch"))

	//r.FindStringSubmatch()
	fmt.Println(r.FindStringSubmatch("peach punch"))

	//r.FindStringSubmatchIndex() 返回匹配的字符串和捕获的字符索引下标
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	//r.FindAllString()，返回所有匹配的字符串，不是只返回匹配到的第一个

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	// 限制匹配的个数，暂定1
	fmt.Println(r.FindAllString("peach punch pinch", 1))
	
	//上述函数都有String，可以通过提供[]bytes参数，然后调用不含有String()的函数达到匹配的目的
	r = regexp.MustCompile("p([a-z]+)ch")
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
