package rawtransactions
//=============== begin:Analyzepsbt ===============//
//合并结合pspt
type Analyzepsbt struct {
	Estimated_vsize float64
	Estimated_feerate float64
	Fee float64
	Next string
	Error string
}
//=============== end:Analyzepsbt===============//

//=============== begin:Combinepsbt ===============//
//合并结合pspt
/*type Combinepsbt struct {

}*/
//=============== end:Combinepsbt===============//

//=============== begin:Combinerawtransaction ===============//
//合并结合原始交易
/*type Combinerawtransaction struct {

}*/
//=============== end:Combinerawtransaction===============//

//=============== begin:Finalizepsbt ===============//
//提取最终的pspt
type Finalizepsbt struct {
	Psbt string
	Hex string
	Complete bool
}
//=============== end:Finalizepsbt===============//

//=============== begin:Createrawtransaction ===============//
//创建一个原始交易
/*type Createrawtransaction struct {

}*/
//=============== end:Createrawtransaction===============//

//=============== begin:Sendrawtransaction ===============//
//找原始交易信息
type Fundrawtransaction struct {
	Result string
	Fee float64
	Changepos float64
}
//=============== end:Fundrawtransaction===============//

//=============== begin: Signrawtransactionwithkey===============//
//用私钥签名交易
type Error struct {
	Txid string
	Vout float64
	ScriptSig string
	Sequence float64
	Error string
}
type  Signrawtransactionwithkey struct {
	Hex string
	Complete bool
	Errors []Error
	Errors_ []interface{}
}
//=============== end:Signrawtransactionwithkey===============//

//=============== begin:Testmempoolaccept ===============//
//测试连接池是否接受链接
type Testing struct {
	Txid string
	Allowed bool
	Reject_reason string
}
type Testmempoolaccept struct {
	Test []Testing
	Test_[]interface{}
}
//=============== end:Testmempoolaccept===============//
