package requests

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

// 枚举, 请求类型
type requestType string

const (
	GET  requestType = "GET"
	POST requestType = "POST"
)

// 生成GET请求
func Get(uri string) *Requests {
	r := newRequest()
	r.reqType = GET
	r.uri = uri
	return r
}

// 生成GET请求
func Post(uri string) *Requests {
	r := newRequest()
	r.reqType = POST
	r.uri = uri
	return r
}

type Requests struct {
	uri string

	// 请求返回的流数据
	resp []byte

	// 请求类型: GET / POST
	reqType requestType

	// 相应建立连接的超时时间
	connTimeout int

	// 相应Resp的超时时间
	respTimeout int

	// 请求代理
	proxy string

	// 请求参数
	param url.Values

	// 认证请求头 basicAuth
	auth *basicAuth

	// 请求头
	header map[string]string

	// Cookie信息
	cookie map[string]string
}

func newRequest() *Requests {
	cookie := make(map[string]string)
	return newRequestWithCookie(cookie)
}

func newRequestWithCookie(cookie map[string]string) *Requests {
	r := &Requests{}

	// 默认请求超时时间
	r.connTimeout = 30_000
	r.respTimeout = 30_000

	r.param = url.Values{}
	r.header = make(map[string]string)
	r.cookie = cookie

	r.reqType = GET
	return r
}

func (r *Requests) dial(netw, addr string) (net.Conn, error) {
	//设置建立连接超时
	conn, err := net.DialTimeout(netw, addr, time.Millisecond*time.Duration(r.connTimeout))
	if err != nil {
		return nil, err
	}
	//设置发送接受数据超时
	_ = conn.SetDeadline(time.Now().Add(time.Millisecond * time.Duration(r.respTimeout)))
	return conn, nil
}

// 建立请求, 并将数据发送给服务器
func (r *Requests) Send() *Requests {

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

	r.setHeader(request)

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("发送请求失败", err.Error())
		return r
	}

	if response == nil {
		return r
	}

	cookies := response.Cookies()
	for _, cookie := range cookies {
		r.AddCookie(cookie.Name, cookie.Value)
	}

	defer response.Body.Close()
	r.resp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取resp流错误", err.Error())
	}

	return r
}

// 将结果请求写入到文件
func (r *Requests) WriteToFile(path string) {
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
func (r *Requests) ReadText() string {
	return string(r.resp)
}
