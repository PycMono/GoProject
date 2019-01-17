package clientUrlBll

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/LoginServer/src/bll/configBDBll/bll"
	. "moqikaka.com/LoginServer/src/model/configDBOrder"
	"moqikaka.com/LoginServer/src/model/login"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/webUtil"
	"testing"
)

func TestName(t *testing.T) {
	url := fmt.Sprintf("%s/%s", "https://managecentertest-slg.moqiplayer.com", "API/ServerList_GetUrl.ashx")
	requestParam := fmt.Sprintf("PartnerID=%d&GameVersionID=%d", 1005, 105)
	webResponse, err := webUtil.PostByteData(url, []byte(requestParam), nil)
	if err != nil {
		fmt.Println(err)
		logUtil.ErrorLog(fmt.Sprintf("con_PackageName=%s,methodName=%s,连接mc/API/ServerList_GetUrl.ashx报错,请检测错误信息ulr=%s,err=%v", con_PackageName, "dfas", url, err))
	}

	var loginReturnObj login.LoginReturnObject
	err = json.Unmarshal(webResponse, &loginReturnObj)
	if err != nil {
		fmt.Println(err)
		logUtil.ErrorLog(fmt.Sprintf("解压loginReturnObj报错：err=%s", err))
	}

	if loginReturnObj.Code != 0 {

	}

	//var tempMap []map[string]string
	// db数据
	clientUrlMap, err := bll.GetItemToMap(GetClientUrl)
	if err != nil {
		return
	}

	_ = clientUrlMap
	//var clientUrl string
	//tempDict := loginReturnObj.Data.([]interface{})[0].(map[string]interface{})
	//if tempDict["OfficialOrTest"] == "1" {
	//	clientUrl = clientUrlMap["zs"]
	//} else {
	//	clientUrl = tempDict["ServerID"].(string)
	//}

	//for _, value := range tempList {
	//	tempDict := value.(map[string]interface{})
	//
	//
	//	for key, value := range ss {
	//		fmt.Println(fmt.Sprintf("%s,%v", key, value))
	//	}
	//}
	//err = json.Unmarshal([]byte(loginReturnObj.Data), &tempMap)
	//fmt.Println(ttt["ServerName"])
}
