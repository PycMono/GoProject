package dal

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"moqikaka.com/Framework/monitorMgr"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/mysqlUtil"
	"moqikaka.com/LoginServer/src/config"
)

var (
	// mysql数据库对象
	dbObj *gorm.DB
)

// 初始化数据库对象
func init() {
	var err error
	configObj := config.GetDBConfig()
	if dbObj, err = gorm.Open("mysql", configObj.ConnectionString); err != nil {
		panic(fmt.Errorf("初始化数据库失败，错误信息为：%s", err))
	}

	if configObj.MaxOpenConns > 0 && configObj.MaxIdleConns > 0 {
		dbObj.DB().SetMaxOpenConns(configObj.MaxOpenConns)
		dbObj.DB().SetMaxIdleConns(configObj.MaxIdleConns)
	}

	monitorMgr.RegisterMonitorFunc(monitor)
}

// 获取Mysql数据库对象
func GetDB() *gorm.DB {
	return dbObj
}

// 读取数据表的所有数据
// dataList:用户保存数据的
func GetAll(dataList interface{}) error {
	if result := dbObj.Find(dataList); result.Error != nil {
		WriteLog("dal.GetAll", result.Error)
		return result.Error
	}

	return nil
}

// 监控mysql连接是否正常
func monitor() error {
	return mysqlUtil.TestConnection(dbObj.DB())
}

// 记录日志
// funcName:方法名称
// err:错误对象
func WriteLog(funcName string, err error) {
	logUtil.NormalLog(fmt.Sprintf("%s出错，错误信息：%s", funcName, err), logUtil.Error)
}
