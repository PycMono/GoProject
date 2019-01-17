package historyServerBll

import (
	"fmt"
	"moqikaka.com/LoginServer/src/bll/configBDBll"
	"moqikaka.com/LoginServer/src/bll/redisBll"
	. "moqikaka.com/LoginServer/src/model/configDBOrder"
	"strconv"
	"strings"
)

var (
	con_PackageName = "historyServerBll"
)

// 获取历史服务器ID
// 参数：
// partnerID：渠道ID
// userID：用户ID
// 返回值：
// 历史服务器ID集合
func GetHistoryServerIDList(partnerID int, userID string) []int {
	// 获取redis所需key字段
	redisKey := makeRedisKey(partnerID, userID, true)
	if redisKey == "" {
		// 兼容刚切换的时候，用户列表丢失的问题
		redisKey = makeRedisKey(partnerID, userID, false)
	}

	// 获取历史服务器
	historyServers := redisBll.GetToString(redisKey)

	// 转换成int集合
	return converToIntList(historyServers)
}

// 保存历史服务器ID
// 参数：
// partnerID：渠道ID
// serverID：服务器ID
// userID：用户ID
// 返回值：
// 1.操作redis返回字符串
// 2.错误对象
func SaveHistoryServerID(partnerID, serverID int, userID string) (string, error) {
	// 获取redis所需key字段
	redisKey := makeRedisKey(partnerID, userID, true)
	if redisKey == "" {
		// 兼容刚切换的时候，用户列表丢失的问题
		redisKey = makeRedisKey(partnerID, userID, false)
	}

	// 获取历史服务器ID
	historyServerIDs := redisBll.GetToString(redisKey)
	nowServerIDs := appendServerID(historyServerIDs, serverID)
	if historyServerIDs == nowServerIDs {
		// 两次服务器ID相同，玩家为改变过服务器
		return "", nil
	}

	fmt.Println(fmt.Sprintf("redisKey=%s,historyServerIDs=%s,nowServerIDs=%s", redisKey, historyServerIDs, nowServerIDs))
	return redisBll.Set(redisKey, nowServerIDs)
}

//---------------------------------私有方法--------------------------

// 构造redis所需key字段
// 参数：
// partnerID：渠道ID
// userID：用户ID
// isMerge：是否合并渠道
// 返回值：
// key信息
func makeRedisKey(partnerID int, userID string, isMerge bool) string {
	// 混服，合作商ID不参与key
	if configBDBll.GetValueToBool(IsAllMixServer) {
		return fmt.Sprintf("%s", userID)
	}

	if !isMerge {
		return fmt.Sprintf("%d_%s", partnerID, userID)
	}

	key := getMergeParterID(partnerID)
	if key == "" {
		return fmt.Sprintf("%d_%s", partnerID, userID)
	}

	return fmt.Sprintf("%s_%s", key, userID)
}

// 追加serverID
// 参数：
// sourceStr：源服务器字符串
// serverID：当前玩家登录的serverID
// 返回值：
// 1.玩家已登录过的服务器ID
func appendServerID(sourceStr string, serverID int) string {
	if sourceStr == "" {
		// 玩家首次登录
		return strconv.Itoa(serverID)
	}

	// 检查是否和上次登录的相同，如果相同不做修改
	serverIDList := converToIntList(sourceStr)
	if len(serverIDList) > 0 && serverIDList[0] == serverID {
		return sourceStr
	}

	var index int
	// 判断是否已经存在了当前服务器ID，如果存在删除，并且把当前服务器ID放在第一位
	for key, item := range serverIDList {
		if item == serverID {
			index = key
			break
		}
	}

	if index > 0 {
		// 剪切数组
		serverIDList = append(serverIDList[:index], serverIDList[index+1:]...)
	}

	// 把当前serverID添加到第一位
	tempServerIDList := make([]int, 0)
	tempServerIDList = append(tempServerIDList, serverID)
	tempServerIDList = append(tempServerIDList, serverIDList...)

	return converToStr(tempServerIDList)
}

// 获取合并渠道名称
// 参数：
// partnerID：渠道ID
// 返回值：
// 1.渠道名称
func getMergeParterID(partnerID int) string {
	partnerIDDict, err := configBDBll.GetValueToMapList(MergeParterIDs)
	if err != nil {
		return ""
	}

	for key, partnerIDList := range partnerIDDict {
		for _, item := range partnerIDList {
			if partnerID == item {
				return key
			}
		}
	}

	return ""
}

// 字符串转成int数组
// 参数：
// str：转换的字符串
// 返回值：
// 1.int数组
func converToIntList(str string) []int {
	var result []int
	if str == "" {
		return result
	}

	splitList := strings.Split(str, ",")
	for _, value := range splitList {
		serverID, err := strconv.Atoi(value)
		if err != nil {
			continue
		}

		result = append(result, serverID)
	}

	return result
}

// 数组转成字符串
// 参数：
// data：待转换的int数组
// 返回值：
// 1.转换后的字符串
func converToStr(data []int) string {
	tempList := make([]string, 0)
	for _, value := range data {
		tempList = append(tempList, strconv.Itoa(value))
	}

	return strings.Join(tempList, ",")
}
