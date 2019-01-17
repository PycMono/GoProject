package config

import (
	"encoding/json"
	"moqikaka.com/goutil/configUtil"
	"moqikaka.com/goutil/debugUtil"
)

// 基础配置对象
type BaseConfig struct {
	// WebServer监听的地址
	WebServerUrl string

	// gops监控端口
	GopsPort string
}

func (this *BaseConfig) String() string {
	bytes, _ := json.Marshal(this)
	return string(bytes)
}

//----------------------------------------------------------------

var (
	baseConfig *BaseConfig
)

// 初始化基础配置
// 参数：
// config：xml对象
// 返回值：
// 1.错误对象
func initBaseConfig(config *configUtil.XmlConfig) error {
	webServerUrl, err := config.String("root/BaseConfig/WebServerUrl", "")
	if err != nil {
		return err
	}

	gopsPort, err := config.String("root/BaseConfig/GopsPort", "")
	if err != nil {
		return err
	}

	baseConfig = &BaseConfig{
		WebServerUrl:              webServerUrl,
		GopsPort:                  gopsPort,
	}

	debugUtil.Println("BaseConfig:", baseConfig)

	return nil
}

// 获取基础对象
// 参数：
// 无
// 返回值：
// 1.基础配置
func GetBaseConfig() *BaseConfig {
	return baseConfig
}
