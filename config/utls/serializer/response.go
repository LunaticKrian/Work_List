package serializer

// Response
// @Func: 基础序列化容器（响应请求）
type Response[T any] struct {
	Status  int    `json:"status"`
	Data    T      `json:"data"`
	Message string `json:"massage"`
	Error   string `json:"error"`
}

// TokenData
// @Func: token信息
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
