package responseObject

// 中心服务器响应结果状态对象
type ResultStatus struct {
	// 状态值(成功是0，非成功以负数来表示)
	Code int

	// 英文信息
	Message string

	// 中文描述
	Desc string `json:"-"`
}

func newResultStatus(code int, message, desc string) *ResultStatus {
	return &ResultStatus{
		Code: code,
		// 兼容客户端显示
		Message: desc,
		Desc:    desc,
	}
}

// 定义所有的响应结果的状态枚举值
var (
	Success              = newResultStatus(0, "Success", "成功")
	DataError            = newResultStatus(-1, "DataError", "数据错误")
	MethodNotDefined     = newResultStatus(-2, "MethodNotDefined", "方法未定义")
	ParamIsEmpty         = newResultStatus(-3, "ParamIsEmpty", "参数为空")
	ParamNotMatch        = newResultStatus(-4, "ParamNotMatch", "参数不匹配")
	ParamTypeError       = newResultStatus(-5, "ParamTypeError", "参数类型错误")
	OnlySupportPOST      = newResultStatus(-6, "OnlySupportPOST", "只支持POST")
	APINotDefined        = newResultStatus(-7, "APINotDefined", "API未定义")
	APIDataError         = newResultStatus(-8, "APIDataError", "API数据错误")
	APIParamError        = newResultStatus(-9, "APIParamError", "API参数错误")
	InvalidIP            = newResultStatus(-10, "InvalidIP", "IP无效")
	ReloadError          = newResultStatus(-11, "ReloadError", "重新加载出错")
	ConfigError          = newResultStatus(-12, "ConfigError", "配置错误")
	ServerGroupNoExists  = newResultStatus(-13, "ServerGroupNoExists", "服务器组不存在")
	ServerGroupDontValid = newResultStatus(-14, "ServerGroupDontValid", "服务器组无效")
	ServerGroupNoOpen    = newResultStatus(-15, "ServerGroupNoOpen", "服务器未开服")
	ConnectMCFailed      = newResultStatus(-16, "ConnectMCFailed", "连接mc中心服务器报错")
	UnmarshalMCFailed    = newResultStatus(-17, "UnmarshalMCFailed", "解压mc返回数据报错")
	DBModelNotExists     = newResultStatus(-18, "DBModelNotExists", "数据库配置不存在")

	// --------------之后从-31开始
	InputDataError   = newResultStatus(-31, "InputDataError", "输入信息错误")
	LoginCheckFailed = newResultStatus(-32, "LoginCheckFailed", "登录校验失败")
	RedisException   = newResultStatus(-33, "RedisException", "Redis操作异常")
)
