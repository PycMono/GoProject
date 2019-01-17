package configBDBll

import (
	"encoding/json"
	"fmt"
	. "moqikaka.com/LoginServer/src/bll/configBDBll/model"
	"moqikaka.com/LoginServer/src/dal"
	. "moqikaka.com/LoginServer/src/model/configDBOrder"
	"moqikaka.com/goutil/logUtil"
	"strconv"
	"sync"
)

var (
	configDBMap = make(map[string]*ConfigDB)
	mutex       sync.RWMutex
)

// 获取指定的模型数据
// 参数：
// key：指定的键
// 返回值：
// 1.模型配置
func GetItem(key string) *ConfigDB {
	mutex.RLock()
	defer mutex.RUnlock()

	value, exists := configDBMap[key]
	if !exists {
		return nil
	}

	return value
}

// 获取指定的value
// 参数：
// key：指定的键
// 返回值：
// 1.value
func GetValueToMap(key string) (map[string]string, error) {
	methodName := "GetItemToMap"
	configDB := GetItem(key)

	var result map[string]string
	err := json.Unmarshal([]byte(configDB.ConfigValue), &result)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s方法报错value=%s", con_PackageName, methodName, configDB.ConfigValue))
		return nil, err
	}

	return result, nil
}

// 获取指定的value
// 参数：
// key：指定的键
// 返回值：
// 1.value
func GetValueToMapList(key string) (map[string][]int, error) {
	methodName := "GetItemToMapList"
	configDB := GetItem(key)

	var result map[string][]int
	err := json.Unmarshal([]byte(configDB.ConfigValue), &result)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s方法报错value=%s", con_PackageName, methodName, configDB.ConfigValue))
		return nil, err
	}

	return result, nil
}

// 获取指定的value
// 参数：
// key：指定的键
// 返回值：
// 1.value
func GetValueToInt32(key string) int {
	configDB := GetItem(key)
	if configDB == nil {
		return 0
	}

	b, _ := strconv.Atoi(configDB.ConfigValue)
	return b
}

// 获取指定的value
// 参数：
// key：指定的键
// 返回值：
// 1.value
func GetValueToString(key string) string {
	configDB := GetItem(key)
	if configDB == nil {
		return ""
	}

	return configDB.ConfigValue
}

// 获取指定的value
// 参数：
// key：指定的键
// 返回值：
// 1.value
func GetValueToBool(key string) bool {
	configDB := GetItem(key)
	if configDB == nil {
		return false
	}

	b, err := strconv.ParseBool(configDB.ConfigValue)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("数据库中key=%s,ConfigValue=%s转成bool值失败", key, configDB.ConfigValue))
		return false
	}

	return b
}

// 获取登录密钥过期时间,默认返回数据库中数组，如果数据库中没有数据返回0
// 参数：无
// 返回值：
// 1.过期时间(秒为单位)
func GetLoginKeyExpireTime() int {
	return GetValueToInt32(LoginKeyExpireTime)
}

//-----------------------------------------------私有方法----------------------------------------------------

// 加载配置文件
func loadConfig() {
	methodName := "loadConfig"

	mutex.RLock()
	defer mutex.RUnlock()

	var configDBList []*ConfigDB
	if err := dal.GetAll(&configDBList); err != nil {
		// 数据库加载失败直接抛出异常
		panic(fmt.Sprintf("%s_%s,数据库基本配置失败,请检查,错误信息:%v", con_PackageName, methodName, err))
	}

	for _, item := range configDBList {
		configDBMap[item.ConfigKey] = item
	}

	msg := fmt.Sprintf("%s_%s数据库中加载成功", con_PackageName, methodName)
	fmt.Println(msg)
	logUtil.DebugLog(msg)
}
