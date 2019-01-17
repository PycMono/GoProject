package impl

import (
	"encoding/json"
	"fmt"
	. "moqikaka.com/LoginServer/src/api"
	api_model "moqikaka.com/LoginServer/src/api/model"
	"moqikaka.com/LoginServer/src/bll/partnerBll/model"
	. "moqikaka.com/LoginServer/src/model/login"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/webUtil"
)

func init() {
	// 注册方法
	Register("MQHY", NewLoginHandler_mqhy)
}

// 摩奇互娱校验助手
type LoginHandler_mqhy struct {
	// 继承父类接口
	*LoginHandlerBase
}

// 登录校验
// 参数：无
// 返回值：
// 1.登录返回结果对象
func (this *LoginHandler_mqhy) CheckUser() *LoginReturnObject {
	methodName := "LoginHandler_mqhy.CheckUser"

	result := NewLoginReturnObject()

	// 为请求所需参数赋值
	appID := this.GetPartnerDB().AppID
	token := this.GetLoginParamToValue("token")
	requestParam := fmt.Sprintf("AppId=%s&Token=%s", appID, token)
	response, err := webUtil.PostByteData(this.GetLoginVerifyUrl(), []byte(requestParam), nil)
	if err != nil {
		msg := fmt.Sprintf("con_PackageName=%s,methodName=%s,GetLoginVerifyUrl=%s,requestParam=%s远程校验出错,请检测错误信息err=%v",
			con_PackageName, methodName, this.GetLoginVerifyUrl(), requestParam, err)
		logUtil.ErrorLog(msg)

		result.Code = ConnectRemoteError
		result.Message = msg

		return result
	}

	// 数据解析
	var sdkReturnObj api_model.SDKReturnObject
	err = json.Unmarshal(response, &sdkReturnObj)
	if err != nil {
		msg := fmt.Sprintf("con_PackageName=%s,methodName=%s,GetLoginVerifyUrl=%s,requestParam=%s解析远程数据错误，response=%s,请检测错误信息err=%v",
			con_PackageName, methodName, this.GetLoginVerifyUrl(), requestParam, string(response), err)
		logUtil.ErrorLog(msg)

		result.Code = AnalysisDataError
		result.Message = msg

		return result
	}

	// 校验结果
	if sdkReturnObj.Code != 0 || sdkReturnObj.Data == "" {
		result.Code = CheckFailed
		// 记录日志
		msg := fmt.Sprintf("con_PackageName=%s,methodName=%s,GetLoginVerifyUrl=%s,requestParam=%s验证失败,code=%v,msg=%v",
			con_PackageName, methodName, this.GetLoginVerifyUrl(), requestParam, sdkReturnObj.Code, sdkReturnObj.Message)
		logUtil.ErrorLog(msg)
		result.Message = msg

		return result
	}

	if sdkReturnObj.ExtraData != "" {
		tempList, _ := json.Marshal(sdkReturnObj.ExtraData) // 序列化数据
		result.ExtraData = string(tempList)
	}

	result.Code = SUCCESS
	result.Data = sdkReturnObj.Data

	return result
}

// 创建摩奇互娱校验助手校验助手
// 参数：
// partnerDB：合作商对象
// loginInfo：登录对象
// 返回值：
// 1.摩奇互娱校验助手对象
func NewLoginHandler_mqhy(partnerDB *model.PartnerDB, loginInfo string) ICheckUser {
	return &LoginHandler_mqhy{
		LoginHandlerBase: NewLoginHandlerBase(partnerDB, loginInfo),
	}
}
