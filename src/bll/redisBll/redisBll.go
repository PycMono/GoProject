package redisBll

import (
	"fmt"
	"github.com/go-redis/redis"
	"moqikaka.com/Framework/monitorMgr"
	"moqikaka.com/LoginServer/src/bll/configBDBll"
	. "moqikaka.com/LoginServer/src/model/configDBOrder"
	"strings"
	"time"
)

var (
	// redis客服端对象
	redisClientObj  *redis.ClusterClient
	con_PackageName = "redisBll"
)

func init() {
	monitorMgr.RegisterMonitorFunc(monitor)
}

// 获取redis中的值
// 参数：
// key：键
// 返回值：
// 1.未处理的redis原始数据
func Get(key string) *redis.StringCmd {
	return redisClientObj.Get(key)
}

// 获取int类型的值
// 参数：
// key：键
// 返回值：
// 1.int类型的值
func GetToInt(key string) (int, error) {
	result := Get(key)
	return result.Int()
}

// 获取string类型的值
// 参数：
// key：键
// 返回值：
// 1.string类型的值
func GetToString(key string) string {
	result := Get(key)
	return result.Val()
}

// 保存到redis中
// 参数：
// key：键
// value：待设置的值
// 返回值：
// 1.操作成功字符串，如果成功是"OK"
// 2.错误对象
func Set(key string, value interface{}) (string, error) {
	expireTime := configBDBll.GetLoginKeyExpireTime()
	return redisClientObj.Set(key, value, time.Duration(expireTime)).Result()
}

// 删除redis中的数据
// 参数：
// key：键
// 返回值：
// 1.删除成功标识，删除成功返回1
// 2.错误对象
func Del(key string) (int64, error) {
	return redisClientObj.Del(key).Result()
}

// 初始化Redis
func InitRedis() {
	methodName := "initRedis"

	if redisClientObj != nil { // 检查是否初始化过
		return
	}

	redisReadWriteHosts := configBDBll.GetValueToString(RedisReadWriteHosts)
	redisDBIndex := configBDBll.GetValueToInt32(RedisDBIndex)
	redisPoolSizeMult := configBDBll.GetValueToInt32(RedisPoolSizeMult)
	redisPwd := configBDBll.GetValueToString(RedisPassword)
	if redisPoolSizeMult <= 0 {
		panic(fmt.Sprintf("%s_%s, InitRedis,RedisPoolSizeMult=%d需要大于0", con_PackageName, methodName, redisPoolSizeMult))
	}

	if redisDBIndex < 0 || redisDBIndex > 15 {
		panic(fmt.Sprintf("%s_%s, InitRedis,redisDBIndex=%d配置超出范围（1-15之间）", con_PackageName, methodName, redisDBIndex))
	}

	readWriteHostsArray := strings.Split(redisReadWriteHosts, "|")
	clusterSlots := func() ([]redis.ClusterSlot, error) {
		var nodList []redis.ClusterNode
		for _, value := range readWriteHostsArray {
			nodList = append(nodList, redis.ClusterNode{
				Addr: value,
			})
		}

		slots := []redis.ClusterSlot{
			{
				Nodes: nodList,
			},
		}

		return slots, nil
	}

	var nodList []redis.ClusterNode
	for _, value := range readWriteHostsArray {
		nodList = append(nodList, redis.ClusterNode{
			Addr: value,
		})
	}

	// 创建节点信息
	redisdb := redis.NewClusterClient(&redis.ClusterOptions{
		ClusterSlots:  clusterSlots,
		RouteRandomly: true,
		Password:      redisPwd,
		PoolSize:      redisPoolSizeMult,
	})

	redisdb.Ping()

	// 设置数据库
	pipe := redisdb.Pipeline()
	pipe.Select(redisDBIndex)
	_, err := pipe.Exec()
	if err != nil {
		panic(fmt.Sprintf("%s_%s, 设置Redis数据库索引(范围[0-15])报错err=%v", con_PackageName, methodName, err))
	}

	err = redisdb.ReloadState()
	if err != nil {
		panic(fmt.Sprintf("%s_%s, ReloadState reloads cluster state报错err=%v", con_PackageName, methodName, err))
	}

	redisClientObj = redisdb

	fmt.Println(fmt.Sprintf("%s_%s,redis初始化成功", con_PackageName, methodName))
}

// 监控redis连接是否正常
// 参数：无
// 返回值：
// 1.错误对象
func monitor() error {
	return redisClientObj.Ping().Err()
}
