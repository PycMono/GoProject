package configBDBll

import "time"

var (
	con_PackageName = "configDB.bll"
)

// 重载配置
func ReloadConfig() {
	loadConfig()           // 重新加载配置
	RefreshWhiteUserInfo() // 刷新白名单用户

	// 启动一个线程定时重新加载数据库
	go func() {
		for {
			time.Sleep(time.Minute * 5)

			loadConfig()           // 重新加载配置
			RefreshWhiteUserInfo() // 刷新白名单用户
		}
	}()
}
