/*
 *Author： Xia
 *周日 12月 27 22:06 2020
 */

package wallet

//=============== begin:AbandontranSactionInfo ===============//
type AbandonTransactionInfo struct {
}

//=============== end:AbandontranSactionInfo ===============//

//=============== begin:AbortRescan ===============//
type AbortRescan struct {
}

//=============== end:AbortRescan ===============//

//=============== begin:AddMultisgAddressInfo ===============//
type AddMultisigAddress struct {
	Address      string
	RedeemScript string
	Descriptor   string
}

//=============== end:AddMultisgAddressInfo ===============//

//=============== begin:UpWallet ===============//

type UpWallet struct {
}

//=============== end:UpWallet ===============//

//=============== begin:Bumpfee ===============//
//!!!!!!!!!!!1
type BumpFee struct {
	Psbt    string
	TxId    string
	Origfee int64
	Fee     int64
	Errors  []string //json array类型
	Str     []string //可能为空
}

//=============== end:Bumpfee ===============//

//=============== begin:Createwallet ===============//
type CreateWallet struct {
	Name    string
	Warning string
}

//=============== end:Createwallet ===============//

//=============== begin:Dumpprivkey ===============//
type DumpPrivkey struct {
	Str string //The private key
}

//=============== end:Dumpprivkey ===============//

//=============== begin:AddressInfo ===============//
type AddressInfo struct {
	Address             string
	ScriptPubKey        string
	Ismine              bool
	Solvable            bool
	Desc                string
	Iswatchonly         bool
	Isscript            bool
	Iswitness           bool
	Hex                 string
	Pubkey              string
	Ischange            bool
	Timestamp           float64
	Hdkeypath           string
	Hdseedid            string
	Hdmasterfingerprint string
	Labels              []string
	Labels_             []interface{}
	Embedded            Embedded
	Embedded_           map[string]interface{}
}

type Embedded struct {
	Isscript        bool
	Iswitness       bool
	Witness_version float64
	Witness_program string
	Pubkey          string
	Address         string
	ScriptPubKey    string
}

//=============== end:AddressInfo ===============//

//=============== begin:Balances ===============//
type Balances struct {
	Mine       Mine
	Mine_      map[float64]interface{}
	Watchonly  Watchonly
	Watchonly_ map[float64]interface{}
}

type Mine struct {
	Trusted           float64
	Untrusted_pending float64
	Immature          float64
	Used              float64
}

type Watchonly struct {
	Trusted           float64
	Untrusted_pending float64
	Immature          float64
}

//=============== end:Balances ===============//

//=============== begin:WalletInfo ===============//
type WalletInfo struct {
	Walletname              string
	Balance                 float64
	Unconfirmed_balance     float64
	Keypoololdest           float64
	Private_keys_enabled    bool
	Hdseedid                string
	Txcount                 float64
	Keypoolsize             float64
	Scanning                bool
	Walletversion           float64
	Immature_balance        float64
	Keypoolsize_hd_internal float64
	Paytxfee                float64
	Avoid_reuse             bool
}

//=============== end:WalletInfo ===============//

//=============== begin:Processpsbt ===============//
//author : 何新萍
type Processps struct {
	Psbt     string
	Complete bool
}

//=============== end:Processpsbt ===============//

//=============== begin:创建者和更新者角色 ===============//
//实现创建者和更新者角色方法的参数相关实体
type Input struct {
	TxId     string
	Vout     float64
	Sequence float64
}

//outputs
type Address struct {
	Address float64
}

type Data struct {
	Data string
}

type Obj struct {
	ChangeAddress          string  //optional, default=pool address
	ChangePosition         float64 //optional, default=random
	Change_type            string  //optional, default=set by -changetype
	IncludeWatching        bool    //optional, default=true
	LockUnspents           bool    //optional, default=false
	FeeRate                float64 //optional, default=0
	SubtractFeeFromOutputs []int   //vout_index optional, default=empty array
	Replaceable            bool    //optional, default=wallet default
	Conf_target            float64 //optional, default=fall back to wallet's confirmation target
	Estimate_mode          string  //optional, default=UNSET, must be one of:"UNSET"、"ECONOMICAL"、"CONSERVATIVE"
}

//=============== end:创建者和更新者角色 ===============//

//=============== begin:用密钥对原始交易进行签名 ===============//
//签名所需参数的实体
type PrevTx struct {
	TxId          string
	Vout          float64
	ScriptPubKey  string
	RedeemScript  string
	WitnessScript string
	Amount        float64
}

//签名的结果实体
type Err struct {
	TxId      string
	Vout      float64
	ScriptSig string
	Sequence  float64
	Error     string
}

type SignResult struct {
	Hex      string
	Complete bool
	Errors   []Err
}

//=============== end:用密钥对原始交易进行签名 ===============//

//=============== begin:WalletFlagInfo ===============//
//更改钱包的给定钱包标志的状态结果信息
type WalletFlagInfo struct {
	Flag_name  string
	Flag_state bool
	Warnings   string
}
//=============== end:WalletFlagInfo ===============//
