package rs

import "bytes"

// 增加一个Cookie参数
func (r *Request) AddCookie(key, value string) *Request {
	r.cookie[key] = value
	return r
}

// 增加多个Cookie参数
func (r *Request) AddCookies(param map[string]string) *Request {
	for k, v := range param {
		r.cookie[k] = v
	}
	return r
}

// 删除一个cookie信息
func (r *Request) RemoveCookies(key string) *Request {
	delete(r.cookie, key)
	return r
}

// 格式化所有cookie信息
func (r *Request) cookieFormat() string {
	bb := bytes.Buffer{}
	for k, v := range r.cookie {
		bb.WriteString(" ;")
		bb.WriteString(k)
		bb.WriteString("=")
		bb.WriteString(v)
	}

	return bb.String()[2:]
}
