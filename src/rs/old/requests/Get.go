package requests

import (
	"net/http"
	"strings"
)

// 生成get请求的url
func (r *Requests) getGetUrl() string {
	uri := r.uri
	param := r.param

	// 参数列表为空 直接返回
	if 0 == len(param) {
		return uri
	}

	body := param.Encode()

	markWithEnd := strings.HasSuffix(uri, "?")

	// 问号结尾, 直接返回uri + param
	if markWithEnd {
		return uri + body
	}

	hasMark := strings.Index(uri, "?") != -1
	// 有问号, 但是在中间
	if hasMark {
		return uri + "&" + body
	}

	// 没有问号
	return uri + "?" + body
}

// get请求将生成get请求url
func (r *Requests) getNewRequest() (*http.Request, error) {
	return http.NewRequest(string(r.reqType), r.getGetUrl(), nil)
}
