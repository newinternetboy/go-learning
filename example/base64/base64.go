package main

import b64 "encoding/base64"

import "fmt"

import "os"

func main() {
	data := "abc123!?$*&()'-=@~"
	//go标准编码
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(sEnc)

	//解码,结尾为+号
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(sDec)
	fmt.Println(string(sDec))

	//go兼容URL编码,结尾为-号
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))

	fmt.Println(uEnc)

	uDec, _ := b64.URLEncoding.DecodeString(uEnc)

	fmt.Println(string(uDec))

	//NewEncoder
	src := []byte("hello world!")
	encoder := b64.NewEncoder(b64.StdEncoding, os.Stdout)
	encoder.Write(src)
	encoder.Close()

	//func NewEncoding(encoder string) *Encoding
	//encoder:64字节长度的字符串，用作转换表
	encodeTransTable := "ab-c-d----------------------------------------------------------"
	fmt.Println([]byte("="))
	enc := b64.NewEncoding(encodeTransTable)

	testString := "this is a test string."
	est := enc.EncodeToString([]byte(testString))

	fmt.Println(est)

}
