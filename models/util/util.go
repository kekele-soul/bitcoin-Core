package util

//=============== begin:CreatemultisigNrequired ===============//
//创建多重签名信息

type CreateMultiSig struct {
	Address      string
	RedeemScript string
	Descriptor   string
}

//=============== end:CreatemultisigNrequired===============//

//=============== begin:EstimateSmartFee ===============//
//估算费用

type EstimateSmartFee struct {
	Feerate float64
	Error   string
	Block   float64
}

//=============== end:EstimateSmartFee===============//

//=============== begin:DesCriptorInfo ===============//
//获取描述符信息
type DesCriptorInfo struct {
	Descriptor     string
	Checksum       string
	Isrange        bool
	Issolvable     bool
	Hasprivatekeys bool
}

//=============== end:DesCriptorInfo==============//

//=============== begin:Validateaddress ===============//
//验证地址信息
type ValidateAddressInfo struct {
	Isvalid         bool
	Address         string
	ScriptPubKey    string
	Isscript        bool
	Iswitness       bool
	Witness_version float64
	Witness_program string
}

//=============== end:Validateaddress=======c========//
