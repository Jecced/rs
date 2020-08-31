package rs

import (
	"net/http"
)

// 设置请求头
func (r *Request) setHeader(request *http.Request) {
	if r.auth != nil {
		request.Header.Set("Authorization", r.auth.encode())
	}
}
