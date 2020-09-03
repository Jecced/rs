package rs

import (
	"net/http"
)

// 设置请求头
func (r *Request) setHeader(request *http.Request) {
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
func (r *Request) AddHeader(key, value string) *Request {
	r.header[key] = value
	return r
}

// 增加多个请求参数
func (r *Request) AddHeaders(param map[string]string) *Request {
	for key, value := range param {
		r.header[key] = value
	}
	return r
}

// 设置请求头中的content type
func (r *Request) SetContentType(contentType string) *Request {
	r.header["Content-Type"] = contentType
	return r
}
