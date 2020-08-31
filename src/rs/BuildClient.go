package rs

import (
	"net/http"
	"time"
)

// 生成建立连接的信息
func (r *Request) buildClient() *http.Client {
	client := &http.Client{
		Transport: r.buildTransport(),
	}
	return client
}

// 生成传输方法
func (r *Request) buildTransport() *http.Transport {
	t := &http.Transport{
		Dial:                  r.dial,
		ResponseHeaderTimeout: time.Millisecond * time.Duration(r.respTimeout),
	}

	r.setProxy(t)
	return t
}
