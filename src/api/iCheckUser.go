package api

import (
	. "moqikaka.com/LoginServer/src/model/login"
)

// 校验接口,check user is available and return *LoginReturnObject
type ICheckUser interface {
	// 登录校验
	CheckUser() *LoginReturnObject
}
