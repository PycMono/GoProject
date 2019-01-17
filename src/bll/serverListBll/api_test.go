package serverListBll

import (
	"encoding/json"
	"fmt"
	. "moqikaka.com/LoginServer/src/model/login"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/webUtil"
	"testing"
)

func TestName(t *testing.T) {
	methodName := "TestName"

	requestParam := fmt.Sprintf("PartnerID=%d&GameVersionID=%d&ResourceVersionID=%d&RandNum=%d&EncryptedString=%s", 1005, 100, 100, 73951302, "0bd378e1e0e7212c0a9ad435c6264870")
	url := fmt.Sprintf("%s/%s", "https://managecentertest-slg.moqiplayer.com", "API/ServerList_Client.ashx")
	response, err := webUtil.PostByteData(url, []byte(requestParam), nil)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("con_PackageName=%s,methodName=%s,连接mc/API/ServerList_Client.ashx报错,请检测错误信息ulr=%s,requestParam=%s,err=%v", con_PackageName, methodName, url, requestParam, err))

	}

	logUtil.DebugLog(fmt.Sprintf("con_PackageName=%s,methodName=%s,url=%s,=%s", con_PackageName, methodName, url, string(response)))

	// 解析返回值
	var loginReturnObj LoginReturnObject
	err = json.Unmarshal(response, &loginReturnObj)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s解析返回值mc返回值loginReturnObj报错：err=%s", con_PackageName, methodName, err))
	}

	tempMap := loginReturnObj.Data.(map[string]interface{})
	resourceObj := tempMap["Resource"]
	var isHasNewResource bool
	resourceMap := resourceObj.(map[string]interface{})
	isHasNewResource = resourceMap != nil && len(resourceMap) > 0
	if isHasNewResource {
		// 有新资源
		return
	}

	//fmt.Println(tempMap["Server"])
	bt, _ := json.Marshal(tempMap["Server"])
	var ttttMap []map[string]interface{}
	_ = json.Unmarshal(bt, &ttttMap)
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range ttttMap {
		if true {
			value["ServerState"] = 1

		}
	}

	for _, value := range ttttMap {
		if true {
			fmt.Println(value["ServerState"])
		}
	}

	//fmt.Println(ttttMap)
}
