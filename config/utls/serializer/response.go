package serializer

// Response
// @Func: 基础序列化容器（响应请求）
type Response[T any] struct {
	Status  int    `json:"status"`
	Data    T      `json:"data"`
	Massage string `json:"massage"`
	Error   string `json:"error"`
}
