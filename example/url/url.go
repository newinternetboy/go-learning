package main

import "fmt"

import "net/url"

import "strings"

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	fmt.Println(u)
	fmt.Println(u.Scheme)
	//u.User认证的用户信息
	fmt.Println(u.User)
	//主机
	host := u.Host
	fmt.Println(host)
	h := strings.Split(host, ":")
	fmt.Println(h)

	//k=>v参数解析
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)

}
