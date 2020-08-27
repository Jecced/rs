package rs

import (
	"fmt"
	"github.com/Jecced/rs/src/rs/util"
	"io/ioutil"
	"net/http"
	"os"
)

type Requests interface {
	// 发送
	Send()
	// 代理
	Proxy()
	// 读取返回流为文本
	ReadText()
	// 将返回流输出文本, 写入到文件
	WriteToFile()
	// 读取成byte流
	ReadByte()
	// 设置请求参数
	SetParams()
	// 设置请求头
	SetHeader()
	// 设置超时时间
	SetTimeOut()
	// 设置Cookie
	SetCookie()
}

const (
	GET  = "GET"
	POST = "POST"
)

type CommRequest interface {
	// 发送请求
	Send() *Request
	// 将结果返回为文本字符串
	ReadText() string
	// 将结果写入到本地文件
	WriteToFile(path string)
}

type Request struct {
	uri     string
	url     string
	client  *http.Client
	resp    []byte
	reqType string
}

func NewRequest() *Request {
	r := &Request{}
	r.client = &http.Client{}
	r.reqType = GET
	return r
}

func (r *Request) Send() *Request {
	request, err := http.NewRequest(r.reqType, r.uri, nil)
	if err != nil {
		fmt.Println("生成请求对象错误")
	}

	response, err := r.client.Do(request)
	if err != nil {
		fmt.Println("发送请求失败")
	}

	if response == nil {
		return r
	}
	defer response.Body.Close()
	r.resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取resp流错误")
	}

	return r
}

// 将结果请求写入到文件
func (r *Request) WriteToFile(path string) {
	// 创建父文件夹
	util.MkdirParent(path)
	// 创建文件
	create, err := os.Create(path)
	if err != nil {
		fmt.Println("创建文件", path, err)
		return
	}

	defer create.Close()

	_, err = create.Write(r.resp)
	if err != nil {
		fmt.Println("写入流出错", err)
		return
	}
}

// 将结果请求读取为字符串
func (r *Request) ReadText() string {
	return string(r.resp)
}
