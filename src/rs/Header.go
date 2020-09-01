package rs

import (
	"net/http"
)

// 设置请求头
func (r *Request) setHeader(request *http.Request) {
	if r.auth != nil {
		request.Header.Set("Authorization", r.auth.encode())
	}

	if r.reqType == POST && len(r.param) > 0 {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
}
