package configBDBll

import (
	"fmt"
	. "moqikaka.com/LoginServer/src/model/configDBOrder"
	"moqikaka.com/goutil/logUtil"
	"strings"
	"sync"
)

var (
	partnerIDMap   = make(map[string]map[string]bool)
	ipMap          = make(map[string]bool)
	whiteUserMutex sync.RWMutex
)

// 刷新白名单信息
func RefreshWhiteUserInfo() {
	methodName := "RefreshWhiteUserInfo"

	whiteUserMutex.RLock()
	defer whiteUserMutex.RUnlock()

	partnerIDMap = make(map[string]map[string]bool)
	ipMap = make(map[string]bool)

	whiteUserConfig := GetItem(WhiteUserList)
	if whiteUserConfig.ConfigValue == "" {
		logUtil.DebugLog(fmt.Sprintf("%s_%s,WhiteUserList配置为%s", con_PackageName, methodName, whiteUserConfig.ConfigValue))
		return
	}

	firstSplitList := strings.Split(whiteUserConfig.ConfigValue, ",")
	for _, value := range firstSplitList {
		secondSplitList := strings.Split(value, ":")
		if len(secondSplitList) != 2 {
			// 合作商与玩家信息一定是2个
			continue
		}

		userIDMap := make(map[string]bool)
		userIDMap[secondSplitList[1]] = true
		partnerIDMap[secondSplitList[0]] = userIDMap
	}

	// 切割IP
	ipConfig := GetItem(WhiteIps)
	if ipConfig.ConfigValue == "" {
		logUtil.DebugLog(fmt.Sprintf("WhiteIps配置为null"))
		return
	} else {
		firstSplitList := strings.Split(ipConfig.ConfigValue, ";")
		for _, value := range firstSplitList {
			ipMap[value] = true
		}
	}

	msg := fmt.Sprintf("%s_%s白名单刷新成功", con_PackageName, methodName)
	fmt.Println(msg)
	logUtil.DebugLog(msg)
}

// 检查是否是白名单用户
// 参数：
// ip：玩家IP地址
// partnerID：玩家合作商ID
// userID：玩家userID
// 返回值：
// 1.true:表示白名单用户，false：非白名单用户
func IsWhiteUser(partnerID, ip, userID string) bool {
	whiteUserMutex.RLock()
	defer whiteUserMutex.RUnlock()

	// 验证是否是白名单IP
	if _, exists := ipMap[ip]; exists {
		return exists
	}

	// 验证合作商ID和UserID是否是白名单用户
	userIDMap, exists := partnerIDMap[partnerID]
	if exists {
		if _, exists := userIDMap[userID]; exists {
			return true
		}
	}

	return false
}
