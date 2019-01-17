package bll

import (
	_ "moqikaka.com/LoginServer/src/api/impl"
	_ "moqikaka.com/LoginServer/src/bll/clientUrlBll"
	"moqikaka.com/LoginServer/src/bll/configBDBll"
	_ "moqikaka.com/LoginServer/src/bll/historyServerBll"
	_ "moqikaka.com/LoginServer/src/bll/loginCheckBll"
	"moqikaka.com/LoginServer/src/bll/partnerBll"
	"moqikaka.com/LoginServer/src/bll/redisBll"
	_ "moqikaka.com/LoginServer/src/bll/serverListBll"
)

// 服务器启动，初始化配置
func Start() {
	// 初始化数据库
	configBDBll.ReloadConfig()

	// 初始化redis
	redisBll.InitRedis()

	// 初始化合作商信息
	partnerBll.RefreshPartnerDB()
}
