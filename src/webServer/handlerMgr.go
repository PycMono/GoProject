package webServer

import (
	"fmt"
	. "moqikaka.com/LoginServer/src/model/responseObject"
)

var (
	// 所有对外提供的处理器集合
	handlerMap = make(map[string]*handler)
)

// 注册API
// path：注册的访问路径
// callback：回调方法
// paramNameList：参数名称集合
func RegisterHandler(path string, callback func(*ApiContext) *ResponseObject) {
	// 判断是否已经注册过，避免命名重复
	if _, exists := handlerMap[path]; exists {
		panic(fmt.Sprintf("%s已经被注册过，请重新命名", path))
	}

	handlerMap[path] = newHandler(path, callback)
}

func getHandler(path string) (*handler, bool) {
	handlerObj, exists := handlerMap[path]

	return handlerObj, exists
}
