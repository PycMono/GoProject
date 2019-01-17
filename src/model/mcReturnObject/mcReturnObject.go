package mcReturnObject

// mc返回信息对象
type MCReturnObject struct {
	// 返回错误码
	Code int

	// 错误信息
	Message string

	// 具体数据
	Data interface{}
}
