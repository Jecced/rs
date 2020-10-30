package requests

// 设置超时时间
func (r *Requests) SetTimeOut(time int) *Requests {
	r.SetConnTimeOut(time)
	r.SetRespTimeOut(time)
	return r
}

// 设置建立连接的超时时间
func (r *Requests) SetConnTimeOut(time int) *Requests {
	r.connTimeout = time
	return r
}

// 设置相应请求的超时时间
func (r *Requests) SetRespTimeOut(time int) *Requests {
	r.respTimeout = time
	return r
}
