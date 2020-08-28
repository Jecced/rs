package rs

import (
	"fmt"
	"github.com/Jecced/rs/src/rs/util"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
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
	// 设置超时时间
	SetTimeOut(time uint8)
	// 设置建立连接的超时时间
	SetConnTimeOut(time uint8)
	// 设置相应请求的超时时间
	SetRespTimeOut(time uint8)
}

type Request struct {
	uri    string
	client *http.Client

	// 请求返回的流数据
	resp []byte

	// 请求类型: GET / POST
	reqType string

	// 相应建立连接的超时时间
	connTimeout uint8

	// 相应Resp的超时时间
	respTimeout uint8

	// ResponseHeaderTimeout
	// 相应Resp请求头的的超时时间
	headTimeout uint8
}

func NewRequest() *Request {
	r := &Request{}

	// 默认请求超时时间
	r.connTimeout = 30
	r.respTimeout = 30
	r.headTimeout = 2

	r.client = &http.Client{}
	r.reqType = GET
	return r
}

func (r *Request) dial(netw, addr string) (net.Conn, error) {
	//设置建立连接超时
	conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(r.connTimeout))
	if err != nil {
		return nil, err
	}
	//设置发送接受数据超时
	_ = conn.SetDeadline(time.Now().Add(time.Second * time.Duration(r.respTimeout)))
	return conn, nil
}

// 设置超时时间
func (r *Request) SetTimeOut(time uint8) *Request {
	r.SetConnTimeOut(time)
	r.SetRespTimeOut(time)
	return r
}

// 设置建立连接的超时时间
func (r *Request) SetConnTimeOut(time uint8) *Request {
	r.connTimeout = time
	return r
}

// 设置相应请求的超时时间
func (r *Request) SetRespTimeOut(time uint8) *Request {
	r.respTimeout = time
	return r
}

// 建立请求, 并将数据发送给服务器
func (r *Request) Send() *Request {

	r.client.Transport = &http.Transport{
		Dial:                  r.dial,
		ResponseHeaderTimeout: time.Second * time.Duration(r.headTimeout),
	}

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
