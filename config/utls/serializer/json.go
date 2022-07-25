package serializer

import "encoding/json"

// ToJson
// @Func: 将结构体对象转换为JSON字符串
func ToJson[T any](t T) string {
	bytes, err := json.Marshal(t)
	if err != nil {
		return ""
	}
	return string(bytes)
}
