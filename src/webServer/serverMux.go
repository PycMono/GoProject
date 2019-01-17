package webServer

import (
	"fmt"
	"net/http"
	"time"

	"encoding/json"
	"moqikaka.com/Framework/ipMgr"
	. "moqikaka.com/LoginServer/src/model/responseObject"
	"moqikaka.com/goutil/debugUtil"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/webUtil"
	"strings"
)

// 定义自定义的Mux对象
type selfDefineMux struct {
}

func (mux *selfDefineMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	responseObj := NewResponseObject()
	startTime := time.Now().Unix()
	endTime := time.Now().Unix()

	// 新建API上下文对象
	context := newApiContext(r, w)

	// 应对NLB的监控
	if r.RequestURI == "/" || r.RequestURI == "/favicon.ico" {
		responseResultDefault(context, responseObj)
		return
	}

	// // 判断是否是POST方法
	// if r.Method != "POST" {
	// 	responseResult(w, responseObj.SetResultStatus(OnlySupportPOST))
	// 	return
	// }

	// 在输出结果给客户端之后再来处理日志的记录，以便于可以尽快地返回给客户端
	defer func() {
		// 记录DEBUG日志
		if debugUtil.IsDebug() || responseObj.ResultStatus != Success {
			//parameter := context.GetRequestString()
			result, _ := json.Marshal(responseObj)

			msg := fmt.Sprintf("%s-->IP:%s;返回数据：%s;", r.RequestURI, webUtil.GetRequestIP(r), string(result))
			logUtil.NormalLog(msg, logUtil.Debug)
		}

		// 超过3s则记录日志
		if endTime-startTime > 3 {
		}
	}()

	// 验证IP是否正确
	if debugUtil.IsDebug() == false && ipMgr.IsIpValid(webUtil.GetRequestIP(r)) == false {
		logUtil.NormalLog(fmt.Sprintf("请求的IP：%s无效", webUtil.GetRequestIP(r)), logUtil.Error)
		responseResultDefault(context, responseObj.SetResultStatus(InvalidIP))
		return
	}

	// 指获取前面的一部分
	api := strings.Split(r.RequestURI, "?")[0]
	// 根据路径选择不同的处理方法
	handlerObj, exists := getHandler(api)
	if !exists {
		responseResultDefault(context, responseObj.SetResultStatus(APINotDefined))
		return
	}

	// 调用方法
	responseObj = handlerObj.handlerFunc(context)
	endTime = time.Now().Unix()

	// 输出结果
	responseResult(context, responseObj)
}
