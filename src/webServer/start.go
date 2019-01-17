package webServer

import (
	"fmt"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/LoginServer/src/config"
	"net/http"
	"sync"
)

// 启动服务器
// wg：WaitGroup对象
func Start(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	configObj := config.GetBaseConfig()
	msg := fmt.Sprintf("WebServer begin to listen:%s", configObj.WebServerUrl)
	logUtil.NormalLog(msg, logUtil.Info)
	debugUtil.Println(msg)

	// 启动Web服务器监听
	if err := http.ListenAndServe(configObj.WebServerUrl, new(selfDefineMux)); err != nil {
		panic(fmt.Errorf("ListenAndServe失败，错误信息为：%s", err))
	}
}
