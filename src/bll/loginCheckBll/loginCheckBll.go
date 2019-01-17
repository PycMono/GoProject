package loginCheckBll

import (
	"fmt"
	"moqikaka.com/LoginServer/src/bll/configBDBll"
	"moqikaka.com/LoginServer/src/bll/redisBll"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/stringUtil"
	"time"
)

var (
	con_PackageName = "loginCheckBll"
)

// 生成登录秘钥
// 参数：
// userID：用户ID
// 返回值：
// 1.登录密钥
func MakeLoginKey(userID string) string {
	methodName := "MakeLoginKey"

	key := stringUtil.GetNewGUID()
	expireTime := configBDBll.GetLoginKeyExpireTime()

	// 获取当前时间的时间戳，加上过期秒数，验证时当前时间超过时间戳则视为过期
	timestamp := time.Now().Unix() + int64(expireTime)

	// 返回的loginKey,GUID_TimeStamp
	loginKey := fmt.Sprintf("%s_%d", key, timestamp)

	// 保存到Redis数据库
	redisBll.Set(key, userID)
	logUtil.DebugLog(fmt.Sprintf("%s_%s,生成密钥成功expireTime=%d,loginKey=%s", con_PackageName, methodName, expireTime, loginKey))

	return loginKey
}

// 登录校验
// 参数：
// loginKey：登录密钥
// userID：用户ID
// 返回值：
// 1.true:验证通过,false:验证未通过
func LoginCheck(loginKey, userID string) bool {
	methodName := "LoginCheck"

	redisUserID := redisBll.GetToString(loginKey)
	if redisUserID != userID {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,登录校验失败loginKey=%s,userID=%s,redisUserID=%s", con_PackageName, methodName, loginKey, userID, redisUserID))
		return false
	}

	logUtil.DebugLog(fmt.Sprintf("%s_%s,登录校验成功loginKey=%s,userID=%s,redisUserID=%s", con_PackageName, methodName, loginKey, userID, redisUserID))
	return true
}

// 创建Related,账号互通相关信息,比如玩家同一个设备创建了两个账号A,B，如果A先创建，B账号创建之后用A账号的数据
// 参数：
// partnerID：渠道ID
// userID：用户ID
// relatedValue：Related值
func SaveRelated(partnerID int, userID, relatedValue string) {
	methodName := "SaveRelated"

	fmt.Println(fmt.Sprintf("partnerID=%d,userID=%s,makeRelatedKey=%s,relatedValue=%s", partnerID, userID, makeRelatedKey(partnerID, userID), relatedValue))
	// 保存到Redis数据库
	redisBll.Set(makeRelatedKey(partnerID, userID), relatedValue)

	// todo Allen,记得删除
	logUtil.DebugLog(fmt.Sprintf("%s_%s,保存Related成功UserID=%s,partnerID=%d,relatedValue=%s", con_PackageName, methodName, userID, partnerID, relatedValue))
}

// 获取Related
// 参数：
// partnerID：渠道ID
// userID：用户ID
// 返回值：
// 1.Related信息
func GetRelated(partnerID int, userID string) string {
	methodName := "GetRelated"

	relatedValue := redisBll.GetToString(makeRelatedKey(partnerID, userID))

	// todo Allen,记得删除
	logUtil.DebugLog(fmt.Sprintf("%s_%s,获取Related失败UserID=%s,partnerID=%d,relatedValue=%s", con_PackageName, methodName, userID, partnerID, relatedValue))

	return relatedValue
}

// 删除Related
// 参数：
// partnerID：渠道ID
// userID：用户ID
// 返回值：
// 无
func DelRelated(partnerID int, userID string) {
	methodName := "DelRelated"

	_, err := redisBll.Del(makeRelatedKey(partnerID, userID))
	if err != nil {
		logUtil.ErrorLog(fmt.Sprintf("%s_%s,删除Related失败UserID=%s,partnerID=%d", con_PackageName, methodName, userID, partnerID))
		return
	}

	// todo Allen,记得删除
	logUtil.DebugLog(fmt.Sprintf("%s_%s,登录校验成功UserID=%s,partnerID=%d", con_PackageName, methodName, userID, partnerID))
}

// 创建RelatedKey
// 参数：
// partnerID：渠道ID
// userID：用户ID
// 返回值：
// 1.RelatedKey
func makeRelatedKey(partnerID int, userID string) string {
	return fmt.Sprintf("%s_%d", userID, partnerID)
}
