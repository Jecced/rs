package rs

import (
	"net/http"
	"net/url"
)

// 代理方法
func (r *Request) proxyFun(_ *http.Request) (*url.URL, error) {
	return url.Parse("http://" + r.proxy)
}

// 设置代理
func (r *Request) setProxy(transport *http.Transport) {
	if r.proxy != "" {
		transport.Proxy = r.proxyFun
	}
}

// 设置请求代理地址
func (r *Request) Proxy(proxy string) *Request {
	r.proxy = proxy
	return r
}
