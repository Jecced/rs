package rs

// 增加一个请求参数
func (r *Request) AddParam(key, value string) *Request {
	r.param[key] = value
	return r
}

// 增加多个请求参数
func (r *Request) AddParams(param map[string]string) *Request {
	for k, v := range param {
		r.param[k] = v
	}
	return r
}
