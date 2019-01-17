package api

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/LoginServer/src/bll/partnerBll/model"
	. "moqikaka.com/LoginServer/src/model/login"
	"moqikaka.com/goutil/logUtil"
	"sync"
)

var (
	funcMap         = make(map[string]func(partnerDB *model.PartnerDB, loginInfo string) ICheckUser)
	mutex           sync.Mutex
	con_PackageName = "api"
)

// 注册方法(如果名称重复会panic)
// name:方法名称（唯一标识）
// definition:方法定义
// 返回值：
// 1.无
func Register(name string, definition func(partnerDB *model.PartnerDB, loginInfo string) ICheckUser) {
	methodName := "Register"

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := funcMap[name]; exists {
		panic(fmt.Sprintf("%s.%s注册方法%s已经存在，请重新命名", con_PackageName, methodName, name))
	}

	funcMap[name] = definition
}

// 调用一个方法
// 参数：
// partnerDB：合作商对象
// loginInfo：登录信息
// 返回值：
// 错误对象
func CallOne(partnerDB *model.PartnerDB, loginInfo string) *LoginReturnObject {
	methodName := "CallOne"

	// 获取方法标识
	tempMap := getOtherConfigInfoToMap(partnerDB)
	name, exists := tempMap["LoginHandler"]
	if !exists {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,partnerDB.OtherConfigInfo=%s不存在LoginHandler的值", con_PackageName, methodName, partnerDB.OtherConfigInfo))
		return nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	if item, exists := funcMap[name]; !exists {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,调用方法%s不存在，请检测", con_PackageName, methodName, name))
		return nil
	} else {
		// 调用方法
		iCheckUser := item(partnerDB, loginInfo)
		return iCheckUser.(ICheckUser).CheckUser()
	}
}

// 获取Partner中其它配置信息
// 参数：无
// 返回值：
// 1.解析后的数据
func getOtherConfigInfoToMap(partnerDB *model.PartnerDB) map[string]string {
	methodName := "getOtherConfigInfoToMap"

	result := make(map[string]string)
	if partnerDB.OtherConfigInfo == "" {
		return result
	}

	// 序列化数据
	err := json.Unmarshal([]byte(partnerDB.OtherConfigInfo), &result)
	if err != nil {
		// 记录日志
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,解析合作商管理中其它配置信息OtherConfigInfo=%s报错err=%v", con_PackageName, methodName, partnerDB.OtherConfigInfo, err))

		return result
	}

	return result
}
