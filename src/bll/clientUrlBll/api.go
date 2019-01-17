package clientUrlBll

//import (
//	"encoding/json"
//	"fmt"
//	"moqikaka.com/LoginServer/src/bll/configBDBll"
//	. "moqikaka.com/LoginServer/src/model/configDBOrder"
//	. "moqikaka.com/LoginServer/src/model/login"
//	. "moqikaka.com/LoginServer/src/model/responseObject"
//	"moqikaka.com/LoginServer/src/webServer"
//	"moqikaka.com/goutil/logUtil"
//	"moqikaka.com/goutil/webUtil"
//)
//
//var (
//	con_PackageName = "clientUrlBll"
//)
//
//func init() {
//	webServer.RegisterHandler("/API/bll/clientUrlBll/getClientUrlHandler", getClientUrlHandler)
//}
//
//// 获取mc地址
//// context:Api请求上下文对象
//// 返回值:
//// 服务器的响应对象（错误码）
//func getClientUrlHandler(context *webServer.ApiContext) *ResponseObject {
//	methodName := "getClientUrlHandler"
//	responseObj := NewResponseObject()
//
//	//var requestObj *model.RequestObject
//	//if err := context.Unmarshal(&requestObj); err != nil {
//	//	return responseObj.SetResultStatus(APIDataError)
//	//}
//	//
//	//// 游戏版本号
//	//gameVersionID := requestObj.GameVersionID
//	//if gameVersionID == 0 {
//	//	gameVersionID = requestObj.GameVersion
//	//}
//	//
//	//if gameVersionID == 0 || requestObj.PartnerID == 0 {
//	//	logUtil.ErrorLog(fmt.Sprintf("%s_%s,开始动登录密钥验证PartnerID=%d,GameVersionID=%d,GameVersion=%d", con_PackageName, methodName, requestObj.PartnerID, gameVersionID, requestObj.GameVersion))
//	//
//	//	responseObj.SetResultStatus(InputDataError)
//	//	return responseObj
//	//}
//	//
//	//url := fmt.Sprintf("%s/%s", configBDBll.GetItem(ManageCenterDomain).ConfigValue, "API/ServerList_GetUrl.ashx")
//	//requestParam := fmt.Sprintf("PartnerID=%d&GameVersionID=%d", requestObj.PartnerID, gameVersionID)
//	//webResponse, err := webUtil.PostByteData(url, []byte(requestParam), nil)
//	//if err != nil {
//	//	logUtil.ErrorLog(fmt.Sprintf("%s_%s,连接mc/API/ServerList_GetUrl.ashx报错,请检测错误信息ulr=%s,requestParam=%s,err=%v", con_PackageName, methodName, url, requestParam, err))
//	//	responseObj.SetResultStatus(ConnectMCFailed)
//	//
//	//	return responseObj
//	//}
//	//
//	//// 解析返回值
//	//var loginReturnObj LoginReturnObject
//	//err = json.Unmarshal(webResponse, &loginReturnObj)
//	//if err != nil {
//	//	logUtil.ErrorLog(fmt.Sprintf("%s_%s,解析返回值loginReturnObj报错：err=%s", con_PackageName, methodName, err))
//	//	responseObj.SetResultStatus(UnmarshalMCFailed)
//	//
//	//	return responseObj
//	//}
//	//
//	//if loginReturnObj.Code != 0 {
//	//	responseObj.Code = loginReturnObj.Code
//	//	responseObj.Message = loginReturnObj.Message
//	//	logUtil.ErrorLog(fmt.Sprintf("%s_%s,解析mc返回数据报错requestParam=%s", con_PackageName, methodName, requestParam))
//	//
//	//	return responseObj
//	//}
//	//
//	//clientUrlMap, err := configBDBll.GetValueToMap(GetClientUrl)
//	//if err != nil {
//	//	responseObj.SetResultStatus(DBModelNotExists)
//	//	logUtil.ErrorLog(fmt.Sprintf("%s_%s,获取DB数据库不存在err=%v", con_PackageName, methodName, err))
//	//
//	//	return responseObj
//	//}
//	//
//	//var clientUrl string
//	//tempDict := loginReturnObj.Data.([]interface{})[0].(map[string]interface{})
//	//
//	//// 判断是否是正式服地址
//	//if tempDict["OfficialOrTest"] == "1" {
//	//	clientUrl = clientUrlMap["zs"]
//	//} else {
//	//	clientUrl = tempDict["ServerID"].(string)
//	//}
//	//
//	//responseObj.Data = clientUrl
//
//	return responseObj
//}
