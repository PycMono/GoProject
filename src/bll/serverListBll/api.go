package serverListBll

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/LoginServer/src/api"
	"moqikaka.com/LoginServer/src/bll/configBDBll"
	"moqikaka.com/LoginServer/src/bll/historyServerBll"
	"moqikaka.com/LoginServer/src/bll/loginCheckBll"
	"moqikaka.com/LoginServer/src/bll/partnerBll"
	. "moqikaka.com/LoginServer/src/model/configDBOrder"
	. "moqikaka.com/LoginServer/src/model/login"
	. "moqikaka.com/LoginServer/src/model/responseObject"
	"moqikaka.com/LoginServer/src/webServer"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/webUtil"
	"strconv"
)

const (
	// 老版本枚举
	ResourceHandleTypeOld = 1

	// 新版本枚举
	ResourceHandleTypeNew = 2
)

var (
	con_PackageName = "serverListBll"
)

func init() {
	webServer.RegisterHandler("/API/bll/serverListBll/getClientListHandler", getClientListHandler)
}

// 获取mc地址
// context:Api请求上下文对象
// 返回值:
// 服务器的响应对象（错误码）
func getClientListHandler(context *webServer.ApiContext) *ResponseObject {
	methodName := "getClientListHandler"
	responseObj := NewResponseObject()

	// 获取参数
	partnerID := context.GetFormValueToInt("PartnerID")
	gameVersionID := context.GetFormValueToInt("GameVersionID")
	randNum := context.GetFormValueToInt("RandNum")
	encryptedString := context.GetFormValue("EncryptedString")
	userID := context.GetFormValue("UserID")
	loginInfo := context.GetFormValue("LoginInfo")

	var resourceVersionID int
	var resourceVersionName string

	// 获取db库中的信息
	resourceHandleType := configBDBll.GetValueToInt32(ResourceHandleType)
	if resourceHandleType == ResourceHandleTypeOld {
		resourceVersionID = context.GetFormValueToInt("ResourceVersionID")
	} else {
		resourceVersionName = context.GetFormValue("ResourceVersionName")
	}

	// loginInfo与userId不能同时为空
	// 当loginInfo非空时，userId为空
	// 当loginInfo为空时，userId非空
	// 需要从第三方获取
	if partnerID == 0 || userID == "" && loginInfo == "" {
		responseObj.SetResultStatus(InputDataError)

		return responseObj
	}

	var requestParam string
	if resourceHandleType == ResourceHandleTypeOld {
		requestParam = fmt.Sprintf("PartnerID=%d&GameVersionID=%d&ResourceVersionID=%d&RandNum=%d&EncryptedString=%s", partnerID, gameVersionID, resourceVersionID, randNum, encryptedString)
	} else {
		requestParam = fmt.Sprintf("PartnerID=%d&GameVersionID=%d&ResourceVersionID=%s&RandNum=%d&EncryptedString=%s", partnerID, gameVersionID, resourceVersionName, randNum, encryptedString)
	}

	url := fmt.Sprintf("%s/%s", configBDBll.GetValueToString(ManageCenterDomain), "API/ServerList_Client.ashx")
	response, err := webUtil.PostByteData(url, []byte(requestParam), nil)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,连接mc/API/ServerList_Client.ashx报错,请检测错误信息ulr=%s,requestParam=%s,err=%v", con_PackageName, methodName, url, requestParam, err))
		responseObj.SetResultStatus(ConnectMCFailed)

		return responseObj
	}

	logUtil.DebugLog(fmt.Sprintf("%s_%s,url=%s,=%s", con_PackageName, methodName, url, string(response)))

	// 解析返回值
	var loginReturnObj LoginReturnObject
	err = json.Unmarshal(response, &loginReturnObj)
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,解析返回值mc返回值loginReturnObj报错：err=%s", con_PackageName, methodName, err))
		responseObj.SetResultStatus(UnmarshalMCFailed)

		return responseObj
	}

	if loginReturnObj.Code != 0 {
		responseObj.Code = loginReturnObj.Code
		responseObj.Message = loginReturnObj.Message
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,mc返回数据错误,code=%d,msg=%s,请检测requestParam=%s", con_PackageName, methodName, loginReturnObj.Code, loginReturnObj.Message, requestParam))

		return responseObj
	}

	tempMap := loginReturnObj.Data.(map[string]interface{})
	resourceObj := tempMap["Resource"]
	var isHasNewResource bool
	if resourceHandleType == ResourceHandleTypeOld {
		resourceList := resourceObj.([]interface{})
		isHasNewResource = resourceList != nil && len(resourceList) > 0

	} else {
		resourceMap := resourceObj.(map[string]interface{})
		isHasNewResource = resourceMap != nil && len(resourceMap) > 0
	}

	if isHasNewResource {
		// 有新资源
		responseObj.Code = HaveNewGameVersion
		return responseObj
	}

	// 定义去检测用户返回值对象
	userInfoMap := make(map[string]interface{})
	// 当有loginInfo，去第三方获取userID
	if loginInfo != "" {
		partnerDB := partnerBll.GetItem(partnerID)
		if partnerDB == nil {
			responseObj.Code = PartnerNotExists
			logUtil.ErrorLog(fmt.Sprintf("%s_%s,未找到，当前合作商列表为partnerID=%s,PartnerInfo=%s", con_PackageName, methodName, partnerID, partnerBll.GetPartnerDBInfo()))

			return responseObj
		}

		// CheckUser进行登录校验
		loginReturnObj := api.CallOne(partnerDB, loginInfo)
		if loginReturnObj.Code != 0 {
			responseObj.Code = loginReturnObj.Code
			responseObj.Message = fmt.Sprintf("%s_%s,msg=%s", con_PackageName, methodName, loginReturnObj.Message)

			return responseObj
		}

		// 重新给userID赋值
		userID = loginReturnObj.Data.(string)
		delRelated := true
		if loginReturnObj.ExtraData != "" {
			var tempMap map[string]interface{}
			err = json.Unmarshal([]byte(loginReturnObj.ExtraData), &tempMap)
			if err != nil {
				msg := fmt.Sprintf("%s_%s,解析ExtraData=%s报错,请检测,err=%v",
					con_PackageName, methodName, loginReturnObj.ExtraData, err)
				logUtil.ErrorLog(msg)

				responseObj.Code = AnalysisDataError
				responseObj.Message = msg

				fmt.Println(responseObj)
				return responseObj
			}

			related, exists := tempMap["Related"]
			if exists {
				// 创建related
				loginCheckBll.SaveRelated(partnerDB.PartnerID, userID, related.(string))
				delRelated = false
			}
		}

		// 如果没有Related,则尝试删除redis中的缓存数据
		// 1.避免一直占据缓存
		// 2.避免因为缓存原因导致的数据不一致问题
		if delRelated {
			loginCheckBll.DelRelated(partnerDB.PartnerID, userID)
		}

		// 初次获取时候返回用户ID和充值用的信息
		userInfoMap["UserID"] = loginReturnObj.Data
		userInfoMap["ExtraData"] = loginReturnObj.ExtraData
		// 生成动态登录密钥
		userInfoMap["LoginInfo"] = loginCheckBll.MakeLoginKey(userID)

		logUtil.DebugLog(fmt.Sprintf("%s_%s,loginInfo=%s", con_PackageName, methodName, userInfoMap["LoginInfo"]))
	}

	// 更改聊天地址和重设测试用户的服务器状态
	// 验证白名单
	isWhite := configBDBll.IsWhiteUser(strconv.Itoa(partnerID), context.GetIP(), userID)
	if isWhite {
		serverByte, _ := json.Marshal(tempMap["Server"])
		var serverList []map[string]interface{}
		_ = json.Unmarshal(serverByte, &serverList)
		for _, serverMap := range serverList {
			if true {
				serverMap["ServerState"] = 1
			}
		}

		tempMap["Server"] = serverList
	}

	// 将登录信息放入返回值对象
	tempMap["UserInfo"] = userInfoMap

	// 将玩家历史登录服务器列表加入到服务器列表中返回
	tempMap["ServerHistory"] = historyServerBll.GetHistoryServerIDList(partnerID, userID)

	// 将组装好的data放入到总的返回对象中
	responseObj.Data = tempMap

	return responseObj
}
