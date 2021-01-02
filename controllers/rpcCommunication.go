package controllers

import (
	"bitcoin-Core/bitcoinServices"
	"bitcoin-Core/models/rpc"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"reflect"
)

type RpcCommunication struct {
	beego.Controller
}

func (r RpcCommunication) Post() {
	requestBody := r.Ctx.Request.Body
	bytes, err2 := ioutil.ReadAll(requestBody)
	if err2 != nil {
		fmt.Println(err2)
	}
	var commit rpc.Communication
	err := json.Unmarshal(bytes,&commit)
	if err != nil {
		r.Ctx.WriteString("解析错误!") //测试用,,
		return
		//r.Ctx.ResponseWriter.Write()//实际用
	}

	bc := bitcoinServices.GetBC()

	typeOf := reflect.TypeOf(bc)
	valueOf := reflect.ValueOf(bc)

	methodName, b := typeOf.MethodByName(commit.Command)
	if !b {
		r.Ctx.WriteString("没有该方法!,请重试!") //测试用,,
		return
	}

	args := []reflect.Value{}
	fmt.Println("方法名为:",methodName.Name)
	name := valueOf.MethodByName(methodName.Name)

	fmt.Printf("%T %v\n", name.Call(args)[0], name.Call(args)[0])

	res := fmt.Sprintf("%v", name.Call(args)[0])

	resJson, err := json.Marshal(res)
	if err != nil {
		fmt.Println("反序列化失败!")
		return
	}
	fmt.Println(resJson)
	r.Ctx.ResponseWriter.Write(resJson)//实际用
	//r.Ctx.WriteString("成功实现!")
}
