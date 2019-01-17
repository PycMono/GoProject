package impl

import (
	"encoding/json"
	"fmt"
	"moqikaka.com/LoginServer/src/bll/partnerBll/model"
	. "moqikaka.com/LoginServer/src/model/login"
	"moqikaka.com/goutil/logUtil"
)

var (
	con_PackageName = "api.impl"
)

// 登陆处理器基类，供所有其它平台的处理类继承
type LoginHandlerBase struct {
	// 合作商对象
	partnerDB *model.PartnerDB

	// 合作商的其它配置信息的Dictionary<string, string>形式
	otherConfigInfoDict map[string]string

	// 登录信息
	loginInfo string

	// 登陆信息的Dictionary<string, string>形式
	loginInfoMap map[string]string
}

// 获取合作商信息
// 参数：无
// 返回值：
// 1.合作商对象
func (this *LoginHandlerBase) GetPartnerDB() *model.PartnerDB {
	return this.partnerDB
}

// 获取Partner中其它配置单条信息
// 参数：
// configName：参数名称
// 返回值：
// 1.其它配置单条信息得值
func (this *LoginHandlerBase) GetOtherConfigToValue(configName string) string {
	methodName := "GetOtherConfigToValue"

	if this.otherConfigInfoDict == nil {
		this.otherConfigInfoDict = this.getOtherConfigInfoToMap()
	}

	value, exists := this.otherConfigInfoDict[configName]
	if !exists {
		// 记录日志
		logUtil.ErrorLog(fmt.Sprintf("%s_%s Partner.OtherConfigInfo中不存在configName=%s的值", con_PackageName, methodName, configName))

		return ""
	}

	return value
}

// 获取登录验证地址
// 参数：无
// 返回值：
// 1.验证地址
func (this *LoginHandlerBase) GetLoginVerifyUrl() string {
	return this.GetOtherConfigToValue("LoginVerifyUrl")
}

// 获取Token的地址
// 参数：无
// 返回值：
// 1.Token的地址
func (this *LoginHandlerBase) GetTokenUrl() string {
	return this.GetOtherConfigToValue("GetTokenUrl")
}

// 获取用户的地址
// 参数：无
// 返回值：
// 1.用户的地址
func (this *LoginHandlerBase) GetUserUrl() string {
	return this.GetOtherConfigToValue("GetUserUrl")
}

// 获取单条登陆参数的值
// 参数：
// paramName：登陆参数的名称
// 返回值：
// 1.登录参数值
func (this *LoginHandlerBase) GetLoginParamToValue(paramName string) string {
	methodName := "GetLoginParamToValue"

	if this.loginInfoMap == nil {
		this.loginInfoMap = this.getLoginParamToMap()
	}

	value, exists := this.loginInfoMap[paramName]
	if !exists {
		// 记录日志
		logUtil.ErrorLog(fmt.Sprintf("%s_%s 登录信息loginInfo中不存在paramName=%s的值", con_PackageName, methodName, paramName))

		return ""
	}

	return value
}

// 判断是否存在accessToken字段
// 参数：无
// 返回值：
// 1.true：存在，false：不存在
func (this *LoginHandlerBase) ExistsAccessToken() bool {
	value := this.GetLoginParamToValue("accessToken")

	return value != ""
}

func (this *LoginHandlerBase) CheckUser() *LoginReturnObject {
	return nil
}

// --------------------------------------------------私有方法------------------------------------------------------

// 获取Partner中其它配置信息
// 参数：无
// 返回值：
// 1.解析后的字典数据
func (this *LoginHandlerBase) getOtherConfigInfoToMap() map[string]string {
	methodName := "getOtherConfigInfoToMap"

	result := make(map[string]string)
	if this.partnerDB.OtherConfigInfo == "" {
		return result
	}

	// 序列化数据
	err := json.Unmarshal([]byte(this.partnerDB.OtherConfigInfo), &result)
	if err != nil {
		// 记录日志
		logUtil.ErrorLog(fmt.Sprintf("%s_%s解析合作商管理中其它配置信息报错err=%v", con_PackageName, methodName, err))

		return result
	}

	return result
}

// 获取登录参数
// 参数：无
// 返回值：
// 1.登录参数解析
func (this *LoginHandlerBase) getLoginParamToMap() map[string]string {
	methodName := "getLoginParamToMap"

	result := make(map[string]string)
	fmt.Println(this.loginInfo)

	// 序列化数据
	err := json.Unmarshal([]byte(this.loginInfo), &result)
	if err != nil {
		// 记录日志
		logUtil.ErrorLog(fmt.Sprintf("%s_%s解析登录参数报错loginInfo=%s,err=%v", con_PackageName, methodName, this.loginInfo, err))

		return result
	}

	return result
}

// 新创建登陆处理器基类对象
// 参数：
// partnerDB：合作商对象
// loginInfo：登录对象
// 返回值：
// 1.登录对象
func NewLoginHandlerBase(partnerDB *model.PartnerDB, loginInfo string) *LoginHandlerBase {
	return &LoginHandlerBase{
		partnerDB: partnerDB,
		loginInfo: loginInfo,
	}
}
