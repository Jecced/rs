package requests

// 枚举, 请求类型
type requestType string

const (
	GET  requestType = "GET"
	POST requestType = "POST"
)

// 枚举, 固定请求头
const (
	HeaderContentType   = "Content-Type"
	HeaderAuthorization = "Authorization"
	HeaderCookie        = "Cookie"
)
