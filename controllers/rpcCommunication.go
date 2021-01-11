package controllers

import (
	"bitcoin-Core/bitcoinServices"
	"bitcoin-Core/models/rpc"
	"bitcoin-Core/utils/structToStr"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"reflect"
	"strings"
)

type RpcCommunication struct {
	beego.Controller
}

func (r RpcCommunication) Post() {
	requestBody := r.Ctx.Request.Body
	bytes, err := ioutil.ReadAll(requestBody)
	if err != nil {
		errStr := fmt.Sprintf("读取请求体出错! err: %s<br>", err)
		r.Ctx.WriteString(errStr)
		return
	}
	var commit rpc.Communication
	err = json.Unmarshal(bytes, &commit)
	if err != nil {
		r.Ctx.WriteString("请求体格式错误,反序列化请求体失败!<br>")
		return
	}

	bc := bitcoinServices.GetBC()
	typeOf := reflect.TypeOf(bc)
	valueOf := reflect.ValueOf(bc)

	methodName, b := typeOf.MethodByName(commit.Command)
	if !b {
		err := fmt.Sprintf("no method: %s<br>", strings.ToLower(commit.Command))
		r.Ctx.WriteString(err)
		return
	}

	args := []reflect.Value{}
	if commit.Params != "" && len(commit.Params) > 7 {
		params := strings.Replace(commit.Params, "\"", "", -1)
		params = strings.Replace(commit.Params, " ", "", -1)
		args = append(args, reflect.ValueOf(params))
	}
	name := valueOf.MethodByName(methodName.Name)

	rpcRes := name.Call(args)[0]
	resStr := ""
	if rpcRes.Kind() == reflect.Struct {
		resStr = structToStr.ToStr(rpcRes.Interface())
	} else {
		resStr = fmt.Sprintf("%v,<br>", rpcRes)
	}
	r.Ctx.WriteString(resStr)
}
