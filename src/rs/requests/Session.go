package requests

type Sessions struct {
	// cookie 请求序列化
	cookie map[string]string

	// 认证请求头 basicAuth
	auth *basicAuth
}

func (s *Sessions) commRequest(t requestType, uri string) *Requests {
	r := newRequestWithCookie(s.cookie)
	r.reqType = t
	r.uri = uri

	if s.auth != nil {
		r.auth = s.auth
	}
	return r
}

func (s *Sessions) Get(uri string) *Requests {
	return s.commRequest(GET, uri)
}

func (s *Sessions) Post(uri string) *Requests {
	return s.commRequest(POST, uri)
}

// 对请求设置basic auth信息
func (s *Sessions) BasicAuth(user, password string) *Sessions {
	s.auth = &basicAuth{name: user, pwd: password}
	return s
}

func Session() *Sessions {
	s := &Sessions{}
	s.cookie = make(map[string]string)
	return s
}
