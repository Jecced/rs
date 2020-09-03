package requests

type Sessions struct {
	// cookie 请求序列化
	cookie map[string]string
}

func (s *Sessions) Get(uri string) *Requests {
	r := newRequestWithCookie(s.cookie)
	r.reqType = GET
	r.uri = uri
	return r
}

func (s *Sessions) Post(uri string) *Requests {
	r := newRequestWithCookie(s.cookie)
	r.reqType = POST
	r.uri = uri
	return r
}

func Session() *Sessions {
	s := &Sessions{}
	s.cookie = make(map[string]string)
	return s
}
