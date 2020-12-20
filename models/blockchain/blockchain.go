//写与区块链相关的实体

package blockchain

//区块信息
type BlockInfo struct {
	Hash              string
	Confirmations     float64
	Strippedsize      float64
	Size              float64
	Weight            float64
	Height            float64
	Version           float64
	VersionHex        string
	Merkleroot        string
	Tx                []interface{}
	Time              float64
	Mediantime        float64
	Nonce             float64
	Bits              string
	Difficulty        float64
	Chainwork         string
	NTx               float64
	Previousblockhash string
}

type Bip34 struct {
	Type   string
	Active bool
	Height float64
}

type Bip66 struct {
	Type   string
	Active bool
	Height float64
}

type Bip65 struct {
	Type   string
	Active bool
	Height float64
}

type Csv struct {
	Type   string
	Active bool
	Height float64
}

type Segwit struct {
	Type   string
	Active bool
	Height float64
}

type Softforks struct {
	Bip65   Bip65
	Bip65_  map[string]interface{}
	Csv     Csv
	Csv_    map[string]interface{}
	Segwit  Segwit
	Segwit_ map[string]interface{}
	Bip34   Bip34
	Bip34_  map[string]interface{}
	Bip66   Bip66
	Bip66_  map[string]interface{}
}

type BlockChainInfo struct {
	Mediantime           float64
	Pruneheight          float64
	Automatic_pruning    bool
	Bestblockhash        string
	Difficulty           float64
	Initialblockdownload bool
	Chainwork            string
	Pruned               bool
	Prune_target_size    float64
	Softforks            Softforks
	Softforks_           map[string]interface{}
	Chain                string
	Warnings             string
	Blocks               float64
	Verificationprogress float64
	Size_on_disk         float64
	Headers              float64
}