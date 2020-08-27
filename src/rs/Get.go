package rs

import "fmt"

func Print() {
	fmt.Println("Test Print, Welcome Use RS Libs")
}

func Test() {
	fmt.Println("Test func")
}

// 生成GET请求
func Get(uri string) *Request {
	r := NewRequest()
	r.reqType = GET
	r.uri = uri
	return r
}

// 生成GET请求
func Post(uri string) *Request {
	r := NewRequest()
	r.reqType = POST
	r.uri = uri
	return r
}
