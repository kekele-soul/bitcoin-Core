package Rpc

import (
	"bitcoin-Core/models/rpc"
	"bitcoin-Core/utils"
	"encoding/json"
	"bitcoin-Core/utils/Base"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

//准备RPC通信标椎格JSON式数据
func PrepareJSON(method string, params []interface{}) string {
	rpcNorm := rpc.RPCNorm{
		Id:      time.Now().Unix(),
		Jsonrpc: utils.RPCVERSION,
		Method:  method,
		Params:  params,
	}

	rpcNormJson, err := json.Marshal(rpcNorm)
	if err != nil {
		return ""
	}
	return string(rpcNormJson)
}

//执行POST请求
func DoPost(url string, header map[string]string, body io.Reader) *rpc.RPCResult {
	client := http.Client{}
	rpcResult := &rpc.RPCResult{}

	//新建一个请求
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		rpcResult.Msg = "请求失败!"
		rpcResult.Data = nil
		return rpcResult
	}

	//设置请求头
	if header != nil {
		for key, value := range header {
			request.Header.Add(key, value)
		}
	}

	//发送请求
	response, err := client.Do(request)
	if err != nil {
		rpcResult.Msg = "请求失败!"
		rpcResult.Data = nil
		return rpcResult
	}

	code := response.StatusCode

	if code == http.StatusOK {
		rpcResult.Code = code
		rpcResult.Msg = "请求成功!"
		responseBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil
		}

		//接收Bitcoin Core 响应回来的结果实体
		res := &rpc.Result{}

		//反序列化
		err = json.Unmarshal(responseBytes, &res)
		if err != nil {
			return nil
		}

		rpcResult.Data = res
		return rpcResult
	} else {
		rpcResult.Code = code
		rpcResult.Msg = "请求失败!"
		rpcResult.Data = nil
	}
	return rpcResult
}


//设置请求头
func RequestHeaders() map[string]string {
	header := make(map[string]string)
	header["Encoding"] = "UTF-8"
	header["Content-Type"] = "application/json"
	header["Authorization"] = "Basic " + Base.Base64Str(utils.RPCUSER + ":" + utils.RPCPASSWORD)
	return header
}