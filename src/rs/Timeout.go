package rs

// 设置超时时间
func (r *Request) SetTimeOut(time int) *Request {
	r.SetConnTimeOut(time)
	r.SetRespTimeOut(time)
	return r
}

// 设置建立连接的超时时间
func (r *Request) SetConnTimeOut(time int) *Request {
	r.connTimeout = time
	return r
}

// 设置相应请求的超时时间
func (r *Request) SetRespTimeOut(time int) *Request {
	r.respTimeout = time
	return r
}
