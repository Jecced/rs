package requests

type Sessions struct {
	// cookie 请求序列化
	cookie map[string]string

	// 认证请求头 basicAuth
	auth *basicAuth

	// 代理
	proxy string

	// 请求头
	header map[string]string
}

func Session() *Sessions {
	s := &Sessions{}
	s.header = make(map[string]string)
	s.cookie = make(map[string]string)
	return s
}

// 生成通用网络请求
func (s *Sessions) commRequest(t requestType, uri string) *Requests {
	r := newRequestWithCookie(s.cookie)
	r.reqType = t
	r.uri = uri

	// 设置通用基础认证信息
	if s.auth != nil {
		r.auth = s.auth
	}

	// 设置通用网络代理信息
	if s.proxy != "" {
		r.proxy = s.proxy
	}

	// 设置通用请求头
	r.AddHeaders(s.header)
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

// 设置通用代理
func (s *Sessions) Proxy(proxy string) *Sessions {
	s.proxy = proxy
	return s
}

// 增加一个请求头参数
func (s *Sessions) AddHeader(key, value string) *Sessions {
	s.header[key] = value
	return s
}

// 增加多个请求头参数
func (s *Sessions) AddHeaders(param map[string]string) *Sessions {
	for key, value := range param {
		s.header[key] = value
	}
	return s
}
