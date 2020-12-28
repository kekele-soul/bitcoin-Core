/*
 *Author： Dudu
 *周二 12月 22 20:50 2020
 */

//写与区块链相关的实体
package blockChain

//=============== begin:BlockInfo ===============//
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

//=============== end:BlockInfo ===============//

//=============== begin:BlockChainInfo ===============//
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

//=============== end:BlockChainInfo ===============//

//=============== begin:Blockstats ===============//
type BlockStats struct {
	Mediantime           float64
	Avgfeerate           float64
	Avgtxsize            float64
	Avgfee               float64
	Swtotal_size         float64
	Txs                  float64
	Ins                  float64
	Minfeerate           float64
	Total_size           float64
	Mediantxsize         float64
	Medianfee            float64
	Minfee               float64
	Swtotal_weight       float64
	Swtxs                float64
	Time                 float64
	Utxo_size_inc        float64
	Blockhash            string
	Maxfeerate           float64
	Maxtxsize            float64
	Mintxsize            float64
	Total_weight         float64
	Totalfee             float64
	Utxo_increase        float64
	Feerate_percentiles  []float64
	Feerate_percentiles_ []interface{}
	Maxfee               float64
	Outs                 float64
	Subsidy              float64
	Total_out            float64
	Height               float64
}

//=============== end:Blockstats ===============//

//=============== begin:NetWorkInfo ===============//
type NetWork struct {
	Name                        string
	Limited                     bool
	Reachable                   bool
	Proxy                       string
	Proxy_randomize_credentials bool
}

type NetWorkInfo struct {
	Localrelay         bool
	Networkactive      bool
	Connections        float64
	Networks           []NetWork
	Networks_          []interface{}
	Version            float64
	Localservices      string
	Warnings           string
	Protocolversion    float64
	Incrementalfee     float64
	Localaddresses     []interface{}
	Subversion         string
	Localservicesnames []interface{}
	Timeoffset         float64
	Relayfee           float64
}

//=============== end:NetWorkInfo ===============//

//=============== begin:BlockHeadInfo ===============//
type BlockHeaderInfo struct {
	Hash              string
	Nonce             float64
	NTx               float64
	Previousblockhash string
	Mediantime        float64
	Version           float64
	VersionHex        string
	Time              float64
	Bits              string
	Difficulty        float64
	Chainwork         string
	Confirmations     float64
	Height            float64
	Merkleroot        string
}

//=============== end:BlockHeadInfo ===============//

//=============== begin:ChainTips ===============//
type Tip struct {
	Height    float64
	Hash      string
	Branchlen float64
	Status    string
}

//区块链Tips信息
type ChainTips struct {
	Tips  []Tip
	Tips_ []interface{}
}

//=============== end:ChainTips ===============//

//=============== begin:ChainTxStats ===============//
//区块链的交易状态信息
type ChainTxStats struct {
	Txrate                    float64
	Time                      float64
	Txcount                   float64
	Window_final_block_hash   string
	Window_final_block_height float64
	Window_block_count        float64
	Window_tx_count           float64
	Window_interval           float64
}

//=============== end:ChainTxStats ===============//

//=============== begin:MempoolAncestorsInfo ===============/
//内存池的祖先信息
type MempoolAncestorsInfo struct {
}

//=============== end:MempoolAncestorsInfo ===============//

//=============== begin:MempoolDescendantsInfo ===============/
//内存池的祖先信息
type MempoolDescendantsInfo struct {
}

//=============== end:MempoolDescendantsInfo ===============//

//=============== begin:MempoolEntryInfo ===============/
//内存池的祖先信息
type MempoolEntryInfo struct {
}

//=============== end:MempoolEntryInfo ===============//

//=============== begin:MempoolInfo ===============/
//获取回收内存信息
type MempoolInfo struct {
	Minrelaytxfee float64
	Loaded        bool
	Size          float64
	Bytes         float64
	Usage         float64
	Maxmempool    float64
	Mempoolminfee float64
}

//=============== end:MempoolInfo ===============//