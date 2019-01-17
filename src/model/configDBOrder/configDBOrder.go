package configDBOrder

const (
	// 聊天服务中心器获取聊天服务器的地址
	ChatServerCenterUrl = "ChatServerCenterUrl"

	// 聊天服务器地址(优先使用配置的地址,未配置时去中心服务器获取),格式:[Address1],[Address2],[Address3],...
	ChatServerInfo = "ChatServerInfo"

	// 客户端Web地址配置
	GetClientUrl = "GetClientUrl"

	// 获取合作商列表的时间间隔（单位：分钟）
	GetPartnerListInterval = "GetPartnerListInterval"

	// ManageCenter域名
	ManageCenterDomain = "ManageCenterDomain"

	// 需要合并的渠道ID
	MergeParterIDs = "MergeParterIDs"

	// 动态密钥Redis数据库索引(范围[0-15])
	RedisDBIndex = "RedisDBIndex"

	// Redis是否使用CLIENT命令（使用：true(默认值) 不使用：false）
	RedisIsUseClientCmd = "RedisIsUseClientCmd"

	// 动态密钥Redis数据库连接池乘数,基数为地址数量
	RedisPoolSizeMult = "RedisPoolSizeMult"

	// 动态密钥Redis数据库地址(只读),多个以"|"符合分割,格式:密码@网络地址:端口号
	RedisReadOnlyHosts = "RedisReadOnlyHosts"

	// 动态密钥Redis数据库地址(读写),多个以"|"符合分割,格式:密码@网络地址:端口号
	RedisReadWriteHosts = "RedisReadWriteHosts"

	// 资源处理方式；1、大主宰、犬夜叉、小精灵；2、大主宰台湾、校花、我欲封天及以后项目
	ResourceHandleType = "ResourceHandleType"

	// 白名单列表（维护状态下可进游戏），格式:ip1;ip2
	WhiteIps = "WhiteIps"

	// 白名单列表（维护状态下可进游戏），格式:[PartnerId1]:[UserId1],[UserId2],...||[PartnerId2]:[UserId3],[UserId4],...
	WhiteUserList = "WhiteUserList"

	// 登录密钥过期时间
	LoginKeyExpireTime = "LoginKeyExpireTime"

	// redis密码
	RedisPassword = "RedisPassword"

	// 是否全部是混服
	IsAllMixServer = "IsAllMixServer"
)
