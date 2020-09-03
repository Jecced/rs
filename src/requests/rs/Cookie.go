package rs

import "bytes"

// 增加一个Cookie参数
func (r *Requests) AddCookie(key, value string) *Requests {
	r.cookie[key] = value
	return r
}

// 增加多个Cookie参数
func (r *Requests) AddCookies(param map[string]string) *Requests {
	for k, v := range param {
		r.cookie[k] = v
	}
	return r
}

// 删除一个cookie信息
func (r *Requests) RemoveCookies(key string) *Requests {
	delete(r.cookie, key)
	return r
}

// 格式化所有cookie信息
func (r *Requests) cookieFormat() string {
	if len(r.cookie) == 0 {
		return ""
	}
	bb := bytes.Buffer{}
	for k, v := range r.cookie {
		bb.WriteString(" ;")
		bb.WriteString(k)
		bb.WriteString("=")
		bb.WriteString(v)
	}

	return bb.String()[2:]
}
