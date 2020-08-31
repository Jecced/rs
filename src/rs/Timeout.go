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

// 设置请求头的超时时间
// 已废弃: 请使用 SetRespTimeOut 来替代 SetHeadTimeOut 方法
// Deprecated: use SetRespTimeOut replace SetHeadTimeOut
func (r *Request) SetHeadTimeOut(time int) *Request {
	//r.headTimeout = time
	r.respTimeout = time
	return r
}
