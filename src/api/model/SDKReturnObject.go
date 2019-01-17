package model

// sdk校验返回对象
type SDKReturnObject struct {
	// 额外信息
	ExtraData interface{}

	// 返回错误码
	Code int

	// 错误信息
	Message string

	// 具体数据
	Data interface{}
}
