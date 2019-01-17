package webServer

import (
	"compress/zlib"
	"encoding/json"
	"fmt"

	. "moqikaka.com/LoginServer/src/model/responseObject"
	"moqikaka.com/goutil/logUtil"
	"moqikaka.com/goutil/zlibUtil"
)

func responseResult(context *ApiContext, responseObj *ResponseObject) {
	data, err := json.Marshal(responseObj)
	if err != nil {
		logUtil.NormalLog("序列化数据出错", logUtil.Error)
		return
	}

	// 判断是否压缩数据
	if responseObj.ZipData {
		data, err = zlibUtil.Compress(data, zlib.DefaultCompression)
		if err != nil {
			logUtil.NormalLog(fmt.Sprintf("压缩数据:%s出错，错误信息:%s", string(data), err), logUtil.Error)
			return
		}
	}

	// 设置跨域问题
	context.responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	context.responseWriter.Write(data)
}

func responseResultDefault(context *ApiContext, responseObj *ResponseObject) {
	data, err := json.Marshal(responseObj)
	if err != nil {
		logUtil.NormalLog("序列化数据出错", logUtil.Error)
		return
	}

	// 判断是否压缩数据
	if responseObj.ZipData {
		data, err = zlibUtil.Compress(data, zlib.DefaultCompression)
		if err != nil {
			logUtil.NormalLog(fmt.Sprintf("压缩数据:%s出错，错误信息:%s", string(data), err), logUtil.Error)
			return
		}
	}

	context.responseWriter.Write(data)
}
