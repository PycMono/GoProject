package login

import "moqikaka.com/LoginServer/src/model/mcReturnObject"

// 登录返回结果对象
type LoginReturnObject struct {
	// 额外信息
	ExtraData string

	// mc返回数据
	mcReturnObject.MCReturnObject
}

// 创建新的登录返回对象
// 参数：无
// 返回值：
// 1.登录返回结果对象
func NewLoginReturnObject() *LoginReturnObject {
	return &LoginReturnObject{}
}
