package rs

import (
	"fmt"
	"github.com/Jecced/rs/src/rs/util"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
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

type Request struct {
	uri string

	// 请求返回的流数据
	resp []byte

	// 请求类型: GET / POST
	reqType string

	// 相应建立连接的超时时间
	connTimeout int

	// 相应Resp的超时时间
	respTimeout int

	// ResponseHeaderTimeout
	// 相应Resp请求头的的超时时间
	headTimeout int

	// 请求代理
	proxy string

	// 请求参数
	param url.Values
}

func NewRequest() *Request {
	r := &Request{}

	// 默认请求超时时间
	r.connTimeout = 30
	r.respTimeout = 30
	r.headTimeout = 2

	r.param = url.Values{}

	r.reqType = GET
	return r
}

func (r *Request) dial(netw, addr string) (net.Conn, error) {
	//设置建立连接超时
	conn, err := net.DialTimeout(netw, addr, time.Millisecond*time.Duration(r.connTimeout))
	if err != nil {
		return nil, err
	}
	//设置发送接受数据超时
	_ = conn.SetDeadline(time.Now().Add(time.Millisecond * time.Duration(r.respTimeout)))
	return conn, nil
}

// 设置超时时间
func (r *Request) SetTimeOut(time int) *Request {
	r.SetConnTimeOut(time)
	r.SetRespTimeOut(time)
	return r
}

// 设置建立连接的超时时间
func (r *Request) SetConnTimeOut(time int) *Request {
	r.connTimeout = time
	return r
}

// 设置相应请求的超时时间
func (r *Request) SetRespTimeOut(time int) *Request {
	r.respTimeout = time
	return r
}

// 设置请求代理
func (r *Request) Proxy(proxy string) *Request {
	r.proxy = proxy
	return r
}

// 建立请求, 并将数据发送给服务器
func (r *Request) Send() *Request {

	var request *http.Request
	var err error

	if r.reqType == GET {
		request, err = r.getNewRequest()
	} else if r.reqType == POST {
		request, err = r.postNewRequest()
	}

	if request == nil && err == nil {
		fmt.Println("无效请求")
		return r
	}

	if err != nil {
		fmt.Println("生成请求对象错误", err.Error())
		return r
	}

	client := r.buildClient()

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("发送请求失败", err.Error())
		return r
	}

	if response == nil {
		return r
	}
	defer response.Body.Close()
	r.resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取resp流错误", err.Error())
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
		fmt.Println("创建文件失败", path, err.Error())
		return
	}

	defer create.Close()

	_, err = create.Write(r.resp)
	if err != nil {
		fmt.Println("写入流出错", err.Error())
		return
	}
}

// 将结果请求读取为字符串
func (r *Request) ReadText() string {
	return string(r.resp)
}
