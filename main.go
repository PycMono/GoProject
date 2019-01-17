package main

import (
	"fmt"
	"github.com/google/gops/agent"
	"moqikaka.com/Framework/monitorMgr"
	"moqikaka.com/LoginServer/src/bll"
	"moqikaka.com/LoginServer/src/config"
	"moqikaka.com/LoginServer/src/webServer"
	"moqikaka.com/goutil/logUtil"
	"sync"
)

var (
	wg            sync.WaitGroup
	con_SEPERATOR = "------------------------------------------------------------------------------"
)

func init() {
	// 设置WaitGroup需要等待的数量，只要有一个服务器出现错误都停止服务器
	wg.Add(1)

	// 设置日志文件的存储目录
	logUtil.SetLogPath("LOG")
}

func main() {
	// 先初始化配置
	bll.Start()

	// 启动监控处理程序
	monitorConfig := config.GetMonitorConfig()
	monitorMgr.Start(
		monitorConfig.ServerIp,
		monitorConfig.ServerName,
		monitorConfig.Interval,
	)

	// 启动监控代理程序
	baseConfig := config.GetBaseConfig()
	if err := agent.Listen(&agent.Options{Addr: baseConfig.GopsPort}); err != nil {
		panic(err)
	}

	// 启动web服务器
	go webServer.Start(&wg)

	logUtil.DebugLog("启动服务器成功...")
	fmt.Println("启动服务器成功...")

	wg.Wait()
}
