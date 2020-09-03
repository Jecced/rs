package rs

import (
	"net/http"
	"strings"
)

// 生成post请求
func (r *Requests) postNewRequest() (*http.Request, error) {
	return http.NewRequest(r.reqType, r.uri, strings.NewReader(r.param.Encode()))
}
