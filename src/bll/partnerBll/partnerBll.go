package partnerBll

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/LoginServer/src/bll/configBDBll"
	"moqikaka.com/LoginServer/src/bll/partnerBll/model"
	. "moqikaka.com/LoginServer/src/model/configDBOrder"
	"moqikaka.com/LoginServer/src/model/mcReturnObject"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/webUtil"
	"sync"
	"time"
)

var (
	partnerMap      = make(map[int]*model.PartnerDB)
	parnerDBInfo    string // 此字段方便查找bug
	mutex           sync.RWMutex
	con_PackageName = "partnerBll.bll"
)

// 获取指定渠道信息
// 参数：
// parnerID：渠道ID
// 返回值：
// 1.渠道信息
func GetItem(parnerID int) *model.PartnerDB {
	mutex.RLock()
	defer mutex.RUnlock()

	partnerDB, exists := partnerMap[parnerID]
	if !exists {
		return partnerDB
	}

	return partnerDB
}

// 获取所有的渠道信息
// 参数：无
// 返回值：
// 1.所有渠道信息
func GetPartnerDBInfo() string {
	return parnerDBInfo
}

// 刷新渠道信息
func RefreshPartnerDB() {
	initPartnerDB()

	// 启动一个线程定时刷新合作商信息
	go func() {
		for {
			time.Sleep(time.Minute * time.Duration(configBDBll.GetValueToInt32("GetPartnerListInterval")))
			initPartnerDB()
		}
	}()
}

// ----------------------------------------------------------初始化合作商------------------------------------------------

// 初始化合作商信息
func initPartnerDB() {
	methodName := "initPartnerDB"

	// 组装url
	url := fmt.Sprintf("%s/%s", configBDBll.GetValueToString(ManageCenterDomain), "API/PartnerList.ashx")
	response, err := webUtil.PostByteData(url, nil, nil)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,连接mc服务器报错url=%s,请检测错误信息err=%v", con_PackageName, methodName, url, err))
		return
	}

	var returnObj mcReturnObject.MCReturnObject
	err = json.Unmarshal(response, &returnObj)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,反序列化mc返回数据报错错误信息err=%v,mc返回数据response=%s:", con_PackageName, methodName, err, string(response)))
		return
	}

	if returnObj.Code != 0 {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,url=%s请求合作商列表返回错误：response=%s", con_PackageName, methodName, url, string(response)))
		return
	}

	var parnerDBList []*model.PartnerDB
	if data, ok := returnObj.Data.(string); !ok {
		msg := fmt.Sprintf("%s_%s获取合作商列表出错：returnObj.Data=%v返回的数据不是string类型", con_PackageName, methodName, returnObj.Data)
		logUtil.ErrorLog(msg)
		return
	} else {
		parnerDBInfo = data
		if err = json.Unmarshal([]byte(data), &parnerDBList); err != nil {
			logUtil.ErrorLog(fmt.Sprintf("%s_%s,反序列化数据出错：err=%s", con_PackageName, methodName, err))
			return
		}
	}

	// 遍历放入临时变量中
	tempPartnerMap := make(map[int]*model.PartnerDB)
	for _, item := range parnerDBList {
		tempPartnerMap[item.PartnerID] = item
	}

	mutex.Lock()
	defer mutex.Unlock()

	partnerMap = tempPartnerMap

	msg := fmt.Sprintf("%s_%s,刷新合作商信息成功", con_PackageName, methodName)
	fmt.Println(msg)
	logUtil.DebugLog(msg)
}
