package loginCheckBll

import (
	"fmt"
	. "moqikaka.com/LoginServer/src/model/responseObject"
	"moqikaka.com/LoginServer/src/webServer"
	"moqikaka.com/goutil/logUtil"
	"strconv"
)

func init() {
	webServer.RegisterHandler("/API/bll/loginCheckBll/checkLoginHandler", checkLoginHandler)
	webServer.RegisterHandler("/API/bll/loginCheckBll/getLatedUserID", getLatedUserID)
}

// 检查登录秘钥是否合法
// context:Api请求上下文对象
// 返回值:
// 服务器的响应对象（错误码）
func checkLoginHandler(context *webServer.ApiContext) *ResponseObject {
	methodName := "checkLoginHandler"
	responseObj := NewResponseObject()
	responseObj.ZipData = false

	userID := context.GetFormValue("UserId")
	loginKey := context.GetFormValue("LoginKey")
	logUtil.DebugLog(fmt.Sprintf("%s_%s,开始动登录密钥验证UserID=%s,LoginKey=%s", con_PackageName, methodName, userID, loginKey))

	if userID == "" || loginKey == "" {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,客服端参数错误UserID=%s,LoginKey=%s", con_PackageName, methodName, userID, loginKey))

		responseObj.SetResultStatus(InputDataError)
		return responseObj
	}

	// 检查登录密钥
	if !LoginCheck(loginKey, userID) {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,动登录密钥验证失败UserID=%s,LoginKey=%s", con_PackageName, methodName, userID, loginKey))

		responseObj.SetResultStatus(LoginCheckFailed)
		return responseObj
	}

	return responseObj
}

// 检查登录秘钥是否合法
// context:Api请求上下文对象
// 返回值:
// 服务器的响应对象（错误码）
func getLatedUserID(context *webServer.ApiContext) *ResponseObject {
	methodName := "getLatedUserID"
	responseObj := NewResponseObject()
	responseObj.ZipData = false

	userID := context.GetFormValue("UserId")
	partnerID, err := strconv.Atoi(context.GetFormValue("PartnerId"))
	if err != nil || userID == "" || partnerID == 0 {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,开始动登录密钥验证UserID=%s,PartnerID=%d", con_PackageName, methodName, userID, partnerID))

		responseObj.SetResultStatus(InputDataError)
		return responseObj
	}

	logUtil.DebugLog(fmt.Sprintf("%s_%s,开始动登录密钥验证UserID=%s,PartnerID=%d", con_PackageName, methodName, userID, partnerID))

	// 获取LatedUserID
	latedUserIDList := GetRelated(partnerID, userID)
	responseObj.Data = latedUserIDList

	return responseObj
}
