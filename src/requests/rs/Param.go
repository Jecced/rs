package rs

// 增加一个请求参数
func (r *Requests) AddParam(key, value string) *Requests {
	r.param.Add(key, value)
	return r
}

// 增加多个请求参数
func (r *Requests) AddParams(param map[string]string) *Requests {
	for k, v := range param {
		r.param.Add(k, v)
	}
	return r
}
