package rs

import (
	"net/http"
	"net/url"
)

// 代理方法
func (r *Requests) proxyFun(_ *http.Request) (*url.URL, error) {
	return url.Parse("http://" + r.proxy)
}

// 设置代理
func (r *Requests) setProxy(transport *http.Transport) {
	if r.proxy != "" {
		transport.Proxy = r.proxyFun
	}
}

// 设置请求代理地址
func (r *Requests) Proxy(proxy string) *Requests {
	r.proxy = proxy
	return r
}
