package config

import (
	"encoding/json"
	"moqikaka.com/goutil/configUtil"
	"moqikaka.com/goutil/debugUtil"
)

// 数据库配置对象
type DBConfig struct {
	// 连接字符串
	ConnectionString string

	// 最大开启连接数量
	MaxOpenConns int

	// 最大空闲连接数量
	MaxIdleConns int
}

func (this *DBConfig) String() string {
	bytes, _ := json.Marshal(this)
	return string(bytes)
}

//----------------------------------------------------------------

var (
	dbConfig *DBConfig
)

// 初始化数据库配置
// 参数：
// config：xml对象
// 返回值：
// 1.错误对象
func initDBConfig(config *configUtil.XmlConfig) error {
	connectionString, err := config.String("root/DBConnection/LoginServerDB", "")
	if err != nil {
		return err
	}

	maxOpenConns, err := config.Int("root/DBConnection/LoginServerDB", "MaxOpenConns")
	if err != nil {
		return err
	}

	maxIdleConns, err := config.Int("root/DBConnection/LoginServerDB", "MaxIdleConns")
	if err != nil {
		return err
	}

	dbConfig = &DBConfig{
		ConnectionString: connectionString,
		MaxOpenConns:     maxOpenConns,
		MaxIdleConns:     maxIdleConns,
	}

	debugUtil.Println("DBConfig:", dbConfig)

	return nil
}

func GetDBConfig() *DBConfig {
	return dbConfig
}
