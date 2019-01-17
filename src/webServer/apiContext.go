package webServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"moqikaka.com/goutil/logUtil"
	"net/http"
	"strconv"
	"strings"
)

// Api请求上下文对象
type ApiContext struct {
	// 请求对象
	request *http.Request

	// 应答写对象
	responseWriter http.ResponseWriter

	// 请求数据
	requestBytes []byte

	// 数据是否已经解析数据
	ifDataParsed bool
}

// 转换内容
func (this *ApiContext) parseContent() error {
	fmt.Println("读取Body")
	defer func() {
		this.request.Body.Close()
		this.ifDataParsed = true
	}()

	data, err := ioutil.ReadAll(this.request.Body)
	if err != nil {
		logUtil.NormalLog(fmt.Sprintf("url:%s,读取数据出错，错误信息为：%s", this.request.RequestURI, err), logUtil.Error)
		return err
	}

	this.requestBytes = data

	return nil
}

// 获取客服端IP信息
// 参数：无
// 返回值：
// 1.ip地址
func (this *ApiContext) GetIP() string {
	return strings.Split(this.request.RemoteAddr, ":")[0]
}

// 获取表单内容
// 参数：
// name：表单值
// 返回值：
// 1.客服端值
func (this *ApiContext) GetFormValue(name string) string {
	return this.request.FormValue(name)
}

// 获取表单内容
// 参数：
// name：表单值
// 返回值：
// 1.客服端值
func (this *ApiContext) GetFormValueToInt(name string) int {
	b, err := strconv.Atoi(this.GetFormValue(name))
	if err != nil {
		return 0
	}

	return b
}

// 获取请求字节数据
// 返回值:
// []byte:请求字节数组
func (this *ApiContext) GetRequestBytes() []byte {
	if this.ifDataParsed == false {
		this.parseContent()
	}

	return this.requestBytes
}

// 获取请求字符串数据
// 返回值:
// 请求字符串数据
func (this *ApiContext) GetRequestString() string {
	data := this.GetRequestBytes()
	if data == nil {
		return ""
	} else {
		return string(data)
	}
}

// 反序列化
// obj:反序列化结果数据
// 返回值:
// 错误对象
func (this *ApiContext) Unmarshal(obj interface{}) error {
	data := this.GetRequestBytes()
	if data == nil {
		return errors.New("RequestBytes为空")
	}

	fmt.Println(string(data))
	//反序列化
	if err := json.Unmarshal(data, &obj); err != nil {
		logUtil.NormalLog(fmt.Sprintf("反序列化%s出错，错误信息为：%s", string(data), err), logUtil.Error)
		return err
	}

	return nil
}

// 新建API上下文对象
// _request:请求对象
// _responseWriter:应答写对象
// 返回值:
// *ApiContext:上下文
func newApiContext(_request *http.Request, _responseWriter http.ResponseWriter) *ApiContext {
	return &ApiContext{
		request:        _request,
		responseWriter: _responseWriter,
		ifDataParsed:   false,
	}
}
