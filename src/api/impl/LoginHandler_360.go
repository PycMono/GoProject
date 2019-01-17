package impl

import (
	"encoding/json"
	"fmt"
	. "moqikaka.com/LoginServer/src/api"
	"moqikaka.com/LoginServer/src/bll/partnerBll/model"
	. "moqikaka.com/LoginServer/src/model/login"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/webUtil"
)

func init() {
	// 注册方法
	Register("360", NewLoginHandler_360)
}

// 360校验助手
type LoginHandler_360 struct {
	// 继承父类接口
	*LoginHandlerBase
}

// 登录校验
// 参数：无
// 返回值：
// 1.登录返回结果对象
func (this *LoginHandler_360) CheckUser() *LoginReturnObject {
	methodName := "LoginHandler_360.CheckUser"

	result := NewLoginReturnObject()

	fields := this.GetOtherConfigToValue("fields")
	access_token := ""

	var tempMap map[string]interface{}
	if !this.ExistsAccessToken() {
		// 为请求所需参数赋值
		grant_type := this.GetOtherConfigToValue("grant_type")
		client_id := this.GetOtherConfigToValue("client_id")
		client_secret := this.GetOtherConfigToValue("client_secret")
		redirect_uri := this.GetOtherConfigToValue("redirect_uri")
		code := this.GetLoginParamToValue("authorizationCode")

		//组装请求参数
		requestParam := fmt.Sprintf("grant_type=%s&code=%s&client_id=%s&client_secret=%s&redirect_uri=%s", grant_type, code, client_id, client_secret, redirect_uri)

		// todo Allen记得解决证书问题
		// ServicePointManager.ServerCertificateValidationCallback = ValidateServerCertificate //为了解决安全证书问题加上的，此行代码不能去掉
		response, err := webUtil.PostByteData(this.GetTokenUrl(), []byte(requestParam), nil)
		if err != nil {
			msg := fmt.Sprintf("con_PackageName=%s,methodName=%s,TokenUrl=%s,requestParam=%s远程校验出错,请检测错误信息err=%v",
				con_PackageName, methodName, this.GetTokenUrl(), requestParam, err)
			logUtil.ErrorLog(msg)

			result.Code = ConnectRemoteError
			result.Message = msg

			return result
		}

		// 解析返回结果
		err = json.Unmarshal(response, &tempMap)
		if err != nil {
			msg := fmt.Sprintf("con_PackageName=%s,methodName=%s,TokenUrl=%s,requestParam=%s远程校验出错,请检测错误信息err=%v",
				con_PackageName, methodName, this.GetTokenUrl(), requestParam, err)
			logUtil.ErrorLog(msg)

			result.Code = AnalysisDataError
			result.Message = msg

			return result
		}

		access_token, _ = tempMap["access_token"].(string)
	} else {
		access_token = this.GetLoginParamToValue("accessToken")
		fields += ",avatar"
	}

	if access_token == "" {
		result.Code = CheckFailed

		return result
	}

	// 组装请求参数
	requestParam := fmt.Sprintf("access_token=%s&fields=%s", access_token, fields)

	// 通过GET方式请求数据
	// ServicePointManager.ServerCertificateValidationCallback = ValidateServerCertificate;//为了解决安全证书问题加上的，此行代码不能去掉
	response, err := webUtil.PostByteData(this.GetUserUrl(), []byte(requestParam), nil)
	if err != nil {
		msg := fmt.Sprintf("con_PackageName=%s,methodName=%s,GetUserUrl=%s,requestParam=%s远程校验出错,请检测错误信息err=%v",
			con_PackageName, methodName, this.GetUserUrl(), requestParam, err)
		logUtil.ErrorLog(msg)

		result.Code = ConnectRemoteError
		result.Message = msg

		return result
	}

	// 解析远程返回值
	err = json.Unmarshal(response, &tempMap)
	if err != nil {
		msg := fmt.Sprintf("con_PackageName=%s,methodName=%s,GetUserUrl=%s,requestParam=%s远程校验出错,请检测错误信息err=%v",
			con_PackageName, methodName, this.GetUserUrl(), requestParam, err)
		logUtil.ErrorLog(msg)

		result.Code = AnalysisDataError
		result.Message = msg

		return result
	}

	// 获取UserID
	userID := tempMap["id"].(string)
	if userID == "" {
		result.Code = CheckFailed

		return result
	}

	extraMap := make(map[string]interface{})
	extraMap["AccessToken"] = access_token
	extraMap["UserID"] = userID
	tempList, _ := json.Marshal(extraMap) // 序列化数据

	result.Code = SUCCESS
	result.Data = userID
	result.ExtraData = string(tempList)

	return result
}

// 创建360校验助手校验助手
// 参数：
// partnerDB：合作商对象
// loginInfo：登录对象
// 返回值：
// 1.360校验助手对象
func NewLoginHandler_360(partnerDB *model.PartnerDB, loginInfo string) ICheckUser {
	return &LoginHandler_360{
		LoginHandlerBase: NewLoginHandlerBase(partnerDB, loginInfo),
	}
}
