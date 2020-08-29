package rs

import (
	"encoding/json"
	"net/http"
	"strings"
)

// get请求将生成get请求url
func (r *Request) postNewRequest() (*http.Request, error) {
	info, _ := json.Marshal(r.param)
	return http.NewRequest(r.reqType, r.uri, strings.NewReader(string(info)))
}
