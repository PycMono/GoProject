package config

import (
	"moqikaka.com/Framework/reloadMgr"
	"moqikaka.com/goutil/configUtil"
	"moqikaka.com/goutil/debugUtil"
)

var (
	// 初始化方法集合
	initFuncList = make([]func(*configUtil.XmlConfig) error, 0, 8)
)

// 初始化数据
func init() {
	registerInitFunc := func(initFunc func(*configUtil.XmlConfig) error) {
		initFuncList = append(initFuncList, initFunc)
	}

	// 注册所有的配置的初始化方法
	registerInitFunc(initBaseConfig)
	registerInitFunc(initDBConfig)
	registerInitFunc(initMonitorConfig)
}

// 初始化数据
func init() {
	if err := reload(); err != nil {
		panic(err)
	}

	reloadMgr.RegisterReloadFunc("config.reload", reload)
}

// 加载项目下面的config文件
// 参数：
// 无
// 返回值：
// 1.错误对象
func reload() error {
	// 读取配置文件内容
	configObj := configUtil.NewXmlConfig()
	err := configObj.LoadFromFile("config.xml")
	if err != nil {
		return err
	}

	debug, err := configObj.Bool("root/DEBUG", "")
	if err != nil {
		return err
	}

	// 设置debugUtil的状态
	debugUtil.SetDebug(debug)

	// 调用所有已经注册的配置初始化方法
	for _, initFunc := range initFuncList {
		if err = initFunc(configObj); err != nil {
			return err
		}
	}

	return nil
}
