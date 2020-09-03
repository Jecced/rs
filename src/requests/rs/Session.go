package rs

type session struct {
	// cookie 请求序列化
	cookie map[string]string
}

func (s *session) Get(uri string) *Requests {
	r := newRequestWithCookie(s.cookie)
	r.reqType = GET
	r.uri = uri
	return r
}

func (s *session) Post(uri string) *Requests {
	r := newRequestWithCookie(s.cookie)
	r.reqType = POST
	r.uri = uri
	return r
}

func Session() *session {
	s := &session{}
	s.cookie = make(map[string]string)
	return s
}
