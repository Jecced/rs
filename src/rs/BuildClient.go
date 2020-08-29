package rs

import (
	"net/http"
	"net/url"
	"time"
)

// 生成建立连接的信息
func (r *Request) buildClient() *http.Client {
	client := &http.Client{}

	t := &http.Transport{
		Dial:                  r.dial,
		ResponseHeaderTimeout: time.Millisecond * time.Duration(r.headTimeout),
	}

	if r.proxy != "" {
		t.Proxy = func(_ *http.Request) (*url.URL, error) {
			return url.Parse("http://" + r.proxy)
		}
	}

	client.Transport = t

	return client
}
