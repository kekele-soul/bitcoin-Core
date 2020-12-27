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
