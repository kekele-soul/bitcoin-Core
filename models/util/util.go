package util

//=============== begin:CreatemultisigNrequired ===============//
//创建多重签名需求

type Createmultisig struct {
	Address string
	RedeemScript string
	Descriptor string
}
//=============== end:CreatemultisigNrequired===============//

//=============== begin:Validateaddress ===============//
//地址起源
type Validateaddress struct {
	Isvalid bool
	Address string
	ScriptPubKey string
	Isscript bool
	Iswitness bool
	Witness_version float64
	Witness_program string
}
//=============== end:Validateaddress=======c========//

//=============== begin:Estimate ===============//
//估算费用

type Estimatesmartfee struct {
	Feerate float64
	Error string
	Block float64
}
//=============== end:Estimatesmartfee===============//

//=============== begin:Getdescriptorinfo ===============//
//获取描述符信息
type Getdescriptorinfo struct {
	Descriptor string
	Checksum string
	Isrange bool
	Issolvable bool
	Hasprivatekeys bool
}
//=============== end:Getdescriptorinfo==============//
