package requests

import (
	"net/http"
)

// 设置请求头
func (r *Requests) setHeader(request *http.Request) {
	if r.auth != nil {
		r.AddHeader("Authorization", r.auth.encode())
	}

	if r.reqType == POST && len(r.param) > 0 {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	request.Header.Set("Cookie", r.cookieFormat())

	for k, v := range r.header {
		request.Header.Set(k, v)
	}
}

// 增加一个请求参数
func (r *Requests) AddHeader(key, value string) *Requests {
	r.header[key] = value
	return r
}

// 增加多个请求参数
func (r *Requests) AddHeaders(param map[string]string) *Requests {
	for key, value := range param {
		r.header[key] = value
	}
	return r
}

// 设置请求头中的content type
func (r *Requests) SetContentType(contentType string) *Requests {
	r.header["Content-Type"] = contentType
	return r
}
