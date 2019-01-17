package responseObject

// Socket服务器的响应对象
type ResponseObject struct {
	// 响应结果的状态值
	*ResultStatus

	// 响应结果的数据
	Data interface{}

	// 是否压缩数据
	ZipData bool
}

func (this *ResponseObject) SetResultStatus(rs *ResultStatus) *ResponseObject {
	this.ResultStatus = rs

	return this
}

func (this *ResponseObject) SetData(data interface{}) *ResponseObject {
	this.Data = data

	return this
}

func NewResponseObject() *ResponseObject {
	return &ResponseObject{
		ResultStatus: Success,
		Data:         nil,
		ZipData:      true,
	}
}
