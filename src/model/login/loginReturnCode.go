package login

const (
	// 成功
	SUCCESS = 0

	// 合作商ID不存在
	PartnerIDNotExists = 1

	// 合作商不存在
	PartnerNotExists = 2

	// 验证失败
	CheckFailed = 3

	// 发生异常
	Exception = 4

	// 远程连接错误
	ConnectRemoteError = 5

	// 有新资源
	HaveNewResource = 6

	// 有新版本
	HaveNewGameVersion = 7

	// 解析远程数据错误
	AnalysisDataError = 8
)
