package requests

import (
	"encoding/base64"
)

// 基础认证
type basicAuth struct {
	name string
	pwd  string
}

func (b *basicAuth) encode() string {
	return "Basic " + base64.URLEncoding.EncodeToString([]byte(b.name+":"+b.pwd))
}

// 对请求设置basic auth信息
func (r *Requests) BasicAuth(user, password string) *Requests {
	r.auth = &basicAuth{name: user, pwd: password}
	return r
}
