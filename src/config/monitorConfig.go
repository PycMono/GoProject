package config

import (
	"encoding/json"
	"moqikaka.com/goutil/configUtil"
	"moqikaka.com/goutil/debugUtil"
)

// 监控配置对象
type MonitorConfig struct {
	// 监控使用的服务器IP
	ServerIp string

	// 监控使用的服务器名称
	ServerName string

	// 监控的时间间隔（单位：分钟）
	Interval int
}

func (this *MonitorConfig) String() string {
	bytes, _ := json.Marshal(this)
	return string(bytes)
}

//----------------------------------------------------------------

var (
	monitorConfig *MonitorConfig
)

// 初始化监控配置
// 参数：
// config：xml对象
// 返回值：
// 1.错误对象
func initMonitorConfig(config *configUtil.XmlConfig) error {
	serverIp, err := config.String("root/MonitorConfig/ServerIp", "")
	if err != nil {
		return err
	}

	serverName, err := config.String("root/MonitorConfig/ServerName", "")
	if err != nil {
		return err
	}

	interval, err := config.Int("root/MonitorConfig/Interval", "")
	if err != nil {
		return err
	}

	monitorConfig = &MonitorConfig{
		ServerIp:   serverIp,
		ServerName: serverName,
		Interval:   interval,
	}

	debugUtil.Println("MonitorConfig:", monitorConfig)

	return nil
}

func GetMonitorConfig() *MonitorConfig {
	return monitorConfig
}
