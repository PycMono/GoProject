package historyServerBll

import (
	"fmt"
	. "moqikaka.com/LoginServer/src/model/responseObject"
	"moqikaka.com/LoginServer/src/webServer"
	"moqikaka.com/goutil/logUtil"
)

func init() {
	webServer.RegisterHandler("/API/bll/historyServerBll/saveServerIDHandler", saveServerIDHandler)
}

// 保存服务器ID
// context:Api请求上下文对象
// 返回值:
// 服务器的响应对象（错误码）
func saveServerIDHandler(context *webServer.ApiContext) *ResponseObject {
	methodName := "getClientUrlHandler"
	responseObj := NewResponseObject()

	partnerID := context.GetFormValueToInt("PartnerID")
	serverID := context.GetFormValueToInt("ServerID")
	userID := context.GetFormValue("UserID")

	if partnerID == 0 || userID == "" || serverID == 0 {
		// 错误打印客服端返回信息
		msg := fmt.Sprintf("%s_%s,保存玩家历史登录服务器,参数错误,请检测PartnerID=%d,ServerID=%d,UserID=%s", con_PackageName,
			methodName, partnerID, serverID, userID)
		logUtil.ErrorLog(msg)

		responseObj.SetResultStatus(InputDataError)
		responseObj.Message = msg
		return responseObj
	}

	// 保存数据
	_, err := SaveHistoryServerID(partnerID, serverID, userID)
	if err != nil {
		responseObj.SetResultStatus(RedisException)
		return responseObj
	}

	return responseObj
}
