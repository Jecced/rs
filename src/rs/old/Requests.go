package old

import "github.com/Jecced/rs/src/rs/old/requests"

// 生成GET请求
func Get(uri string) *requests.Requests {
	return requests.Get(uri)
}

// 生成GET请求
func Post(uri string) *requests.Requests {
	return requests.Post(uri)
}

// 生成持久维护cookie的请求
func Session() *requests.Sessions {
	return requests.Session()
}
