package webServer

import (
	. "moqikaka.com/LoginServer/src/model/responseObject"
)

// 请求方法对象
type handler struct {
	// 注册的访问路径
	path string

	// 方法定义
	handlerFunc func(context *ApiContext) *ResponseObject
}

// 创建新的请求方法对象
func newHandler(_path string, _handlerFunc func(context *ApiContext) *ResponseObject) *handler {
	return &handler{
		path:        _path,
		handlerFunc: _handlerFunc,
	}
}
