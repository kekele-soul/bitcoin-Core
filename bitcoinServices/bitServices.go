//此处写RPC通信方法

package bitcoinServices

import (
	"bitcoin-Core/models/blockchain"
	"bitcoin-Core/models/rawTransactions"
	"bitcoin-Core/models/util"
	"bitcoin-Core/models/wallet"
	"bitcoin-Core/utils"
	"bitcoin-Core/utils/Rpc"
	"strings"
)

type btcSer struct {
	BlockChahin
	Control
	Generating
	Mining
	Network
	Util
	Wallet
	Zmq
}

func GetBC() btcSer {
	return btcSer{}
}

//获取最高区块Hash,成功返回最高区块Hash,否则返回空字符串
func (bc btcSer) GetBestBlockHash() string {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBESTBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//根据区块Hash获取区块信息
func (bc btcSer) GetBlockInfoByHash(blockHash string) blockchain.BlockInfo {
	paramsSlice := []interface{}{blockHash}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCK, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	blockInfo := blockchain.BlockInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockInfo.Time = resBytes["time"].(float64)
		blockInfo.Bits = resBytes["bits"].(string)
		blockInfo.NTx = resBytes["nTx"].(float64)
		blockInfo.Previousblockhash = resBytes["previousblockhash"].(string)
		blockInfo.Strippedsize = resBytes["strippedsize"].(float64)
		blockInfo.Size = resBytes["size"].(float64)
		blockInfo.Merkleroot = resBytes["merkleroot"].(string)
		blockInfo.Weight = resBytes["weight"].(float64)
		blockInfo.Version = resBytes["version"].(float64)
		blockInfo.VersionHex = resBytes["versionHex"].(string)
		blockInfo.Nonce = resBytes["nonce"].(float64)
		blockInfo.Difficulty = resBytes["difficulty"].(float64)
		blockInfo.Hash = resBytes["hash"].(string)
		blockInfo.Height = resBytes["height"].(float64)
		blockInfo.Mediantime = resBytes["mediantime"].(float64)
		blockInfo.Chainwork = resBytes["chainwork"].(string)
		blockInfo.Confirmations = resBytes["confirmations"].(float64)
		blockInfo.Tx = resBytes["tx"].([]interface{})
	}

	return blockInfo
}

//根据区块hash获取区块Info
func (bc btcSer) GetBlockInfoByHeight(height int) blockchain.BlockInfo {
	blockInfo := blockchain.BlockInfo{}
	if float64(height) > bc.GetBlockCount() {
		return blockInfo
	}

	paramsSlice := []interface{}{height}

	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return bc.GetBlockInfoByHash(hash)
	}

	return blockInfo
}

//获取区块链信息
func (bc btcSer) GetBlockChainInfo() blockchain.BlockChainInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKCHAININFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	blockChainInfo := blockchain.BlockChainInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockChainInfo.Size_on_disk = resBytes["size_on_disk"].(float64)
		blockChainInfo.Blocks = resBytes["blocks"].(float64)
		blockChainInfo.Mediantime = resBytes["mediantime"].(float64)
		blockChainInfo.Chainwork = resBytes["chainwork"].(string)
		blockChainInfo.Automatic_pruning = resBytes["automatic_pruning"].(bool)
		blockChainInfo.Pruneheight = resBytes["pruneheight"].(float64)
		blockChainInfo.Bestblockhash = resBytes["bestblockhash"].(string)
		blockChainInfo.Difficulty = resBytes["difficulty"].(float64)
		blockChainInfo.Initialblockdownload = resBytes["initialblockdownload"].(bool)
		blockChainInfo.Pruned = resBytes["pruned"].(bool)
		blockChainInfo.Prune_target_size = resBytes["prune_target_size"].(float64)
		blockChainInfo.Chain = resBytes["chain"].(string)
		blockChainInfo.Warnings = resBytes["warnings"].(string)
		blockChainInfo.Headers = resBytes["headers"].(float64)

		blockChainInfo.Softforks_, ok = resBytes["softforks"].(map[string]interface{})
		if ok {
			blockChainInfo.Softforks.Bip34_, ok = blockChainInfo.Softforks_["bip34"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip34.Height = blockChainInfo.Softforks.Bip34_["height"].(float64)
				blockChainInfo.Softforks.Bip34.Type = blockChainInfo.Softforks.Bip34_["type"].(string)
				blockChainInfo.Softforks.Bip34.Active = blockChainInfo.Softforks.Bip34_["active"].(bool)
			}

			blockChainInfo.Softforks.Bip65_, ok = blockChainInfo.Softforks_["bip65"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip65.Height = blockChainInfo.Softforks.Bip65_["height"].(float64)
				blockChainInfo.Softforks.Bip65.Type = blockChainInfo.Softforks.Bip65_["type"].(string)
				blockChainInfo.Softforks.Bip65.Active = blockChainInfo.Softforks.Bip65_["active"].(bool)
			}

			blockChainInfo.Softforks.Bip66_, ok = blockChainInfo.Softforks_["Bip66"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Bip66.Height = blockChainInfo.Softforks.Bip66_["height"].(float64)
				blockChainInfo.Softforks.Bip66.Type = blockChainInfo.Softforks.Bip66_["type"].(string)
				blockChainInfo.Softforks.Bip66.Active = blockChainInfo.Softforks.Bip66_["active"].(bool)
			}

			blockChainInfo.Softforks.Segwit_, ok = blockChainInfo.Softforks_["Segwit"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Segwit.Height = blockChainInfo.Softforks.Segwit_["height"].(float64)
				blockChainInfo.Softforks.Segwit.Type = blockChainInfo.Softforks.Segwit_["type"].(string)
				blockChainInfo.Softforks.Segwit.Active = blockChainInfo.Softforks.Segwit_["active"].(bool)
			}

			blockChainInfo.Softforks.Csv_, ok = blockChainInfo.Softforks_["Csv"].(map[string]interface{})
			if ok {
				blockChainInfo.Softforks.Csv.Height = blockChainInfo.Softforks.Csv_["height"].(float64)
				blockChainInfo.Softforks.Csv.Type = blockChainInfo.Softforks.Csv_["type"].(string)
				blockChainInfo.Softforks.Csv.Active = blockChainInfo.Softforks.Csv_["active"].(bool)
			}

		}
	}

	return blockChainInfo
}

//获取网络信息
func (bc btcSer) GetNetWorkInfo() blockchain.NetWorkInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETNETWORKINFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	netWorkInfo := blockchain.NetWorkInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		netWorkInfo.Relayfee = resBytes["relayfee"].(float64)
		netWorkInfo.Warnings = resBytes["warnings"].(string)
		netWorkInfo.Localrelay = resBytes["localrelay"].(bool)
		netWorkInfo.Networks_, ok = resBytes["networks"].([]interface{})
		if ok {
			for i := 0; i < len(netWorkInfo.Networks_); i++ {
				mapValue, ok := netWorkInfo.Networks_[i].(map[string]interface{})
				if ok {
					var network blockchain.NetWork
					network.Name = mapValue["name"].(string)
					network.Limited = mapValue["limited"].(bool)
					network.Reachable = mapValue["reachable"].(bool)
					network.Proxy = mapValue["proxy"].(string)
					network.Proxy_randomize_credentials = mapValue["proxy_randomize_credentials"].(bool)
					netWorkInfo.Networks = append(netWorkInfo.Networks, network)
				}
			}
		}

		netWorkInfo.Version = resBytes["version"].(float64)
		netWorkInfo.Subversion = resBytes["subversion"].(string)
		netWorkInfo.Protocolversion = resBytes["protocolversion"].(float64)
		netWorkInfo.Timeoffset = resBytes["timeoffset"].(float64)
		netWorkInfo.Networkactive = resBytes["networkactive"].(bool)
		netWorkInfo.Localaddresses = resBytes["localaddresses"].([]interface{})
		netWorkInfo.Localservices = resBytes["localservices"].(string)
		netWorkInfo.Localservicesnames = resBytes["localservicesnames"].([]interface{})
		netWorkInfo.Connections = resBytes["connections"].(float64)
		netWorkInfo.Incrementalfee = resBytes["incrementalfee"].(float64)
	}

	return netWorkInfo
}

//获取区块总数
func (bc btcSer) GetBlockCount() float64 {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKCOUNT, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -1
}

//根据区块高度获取区块的hash
func (bc btcSer) GetBlockHash(height int) string {
	if float64(height) > bc.GetBlockCount() {
		return ""
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return hash
	}

	return ""
}

//根据区块Hash获取区块头信息
func (bc btcSer) GetBlockHeaderInfoByHash(hash string) blockchain.BlockHeaderInfo {
	paramsSlice := []interface{}{hash}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKHEADER, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列华操作
	blockHeadInfo := blockchain.BlockHeaderInfo{}

	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockHeadInfo.Height = resBytes["height"].(float64)
		blockHeadInfo.Mediantime = resBytes["mediantime"].(float64)
		blockHeadInfo.Chainwork = resBytes["chainwork"].(string)
		blockHeadInfo.NTx = resBytes["nTx"].(float64)
		blockHeadInfo.Previousblockhash = resBytes["previousblockhash"].(string)
		blockHeadInfo.Hash = resBytes["hash"].(string)
		blockHeadInfo.Version = resBytes["version"].(float64)
		blockHeadInfo.Merkleroot = resBytes["merkleroot"].(string)
		blockHeadInfo.Time = resBytes["time"].(float64)
		blockHeadInfo.Difficulty = resBytes["difficulty"].(float64)
		blockHeadInfo.Confirmations = resBytes["confirmations"].(float64)
		blockHeadInfo.VersionHex = resBytes["versionHex"].(string)
		blockHeadInfo.Nonce = resBytes["nonce"].(float64)
		blockHeadInfo.Bits = resBytes["bits"].(string)
	}

	return blockHeadInfo
}

//根据区块Height获取区块头信息
func (bc btcSer) GetBlockHeaderInfoByHeight(height float64) blockchain.BlockHeaderInfo {
	blockHeaderInfo := blockchain.BlockHeaderInfo{}
	if float64(height) > bc.GetBlockCount() {
		return blockHeaderInfo
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	hash, ok := rpcResult.Data.Result.(string)
	if ok {
		return bc.GetBlockHeaderInfoByHash(hash)
	}

	return blockHeaderInfo
}

//根据区块高度获取区块状态
func (bc btcSer) GetBlockStatsInfoByHeight(height float64) blockchain.BlockStats {
	blockStats := blockchain.BlockStats{}
	if height > bc.GetBlockCount() {
		return blockStats
	}

	paramsSlice := []interface{}{height}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON("getblockstats", paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		blockStats.Subsidy = res["subsidy"].(float64)
		blockStats.Total_weight = res["total_weight"].(float64)
		blockStats.Height = res["height"].(float64)
		blockStats.Medianfee = res["medianfee"].(float64)
		blockStats.Mediantxsize = res["mediantxsize"].(float64)
		blockStats.Totalfee = res["totalfee"].(float64)
		blockStats.Avgfee = res["avgfee"].(float64)
		blockStats.Avgfeerate = res["avgfeerate"].(float64)
		blockStats.Minfee = res["minfee"].(float64)
		blockStats.Swtxs = res["swtxs"].(float64)
		blockStats.Time = res["time"].(float64)
		blockStats.Total_size = res["total_size"].(float64)
		blockStats.Blockhash = res["blockhash"].(string)
		blockStats.Feerate_percentiles_, ok = res["feerate_percentiles"].([]interface{})
		if ok {
			for i := 0; i < len(blockStats.Feerate_percentiles_); i++ {
				value, ok := blockStats.Feerate_percentiles_[i].(float64)
				if ok {
					blockStats.Feerate_percentiles = append(blockStats.Feerate_percentiles, value)
				}
			}
		}

		blockStats.Ins = res["ins"].(float64)
		blockStats.Minfeerate = res["minfeerate"].(float64)
		blockStats.Avgtxsize = res["avgtxsize"].(float64)
		blockStats.Maxfeerate = res["maxfeerate"].(float64)
		blockStats.Mintxsize = res["mintxsize"].(float64)
		blockStats.Outs = res["outs"].(float64)
		blockStats.Utxo_increase = res["utxo_increase"].(float64)
		blockStats.Mediantime = res["mediantime"].(float64)
		blockStats.Swtotal_weight = res["swtotal_weight"].(float64)
		blockStats.Txs = res["txs"].(float64)
		blockStats.Utxo_size_inc = res["utxo_size_inc"].(float64)
		blockStats.Maxfee = res["maxfee"].(float64)
		blockStats.Swtotal_size = res["swtotal_size"].(float64)
		blockStats.Total_out = res["total_out"].(float64)
	}

	return blockStats
}

//根据区块Hash获取区块状态
func (bc btcSer) GetBlockStatsInfoByHash(hash string) blockchain.BlockStats {
	blockStats := blockchain.BlockStats{}
	if len(hash) != 64 || hash[0] != 48{
		return blockStats
	}

	var height float64 = -1
	if hash == bc.GetBlockHash(0) {
		return bc.GetBlockStatsInfoByHeight(0)
	}

	blockInfo := bc.GetBlockInfoByHash(hash)
	height = blockInfo.Height

	if height > 0 {
		return bc.GetBlockStatsInfoByHeight(height)
	}

	return blockStats
}

//获取区块链Tip信息
func (bc btcSer) GetChainTips() blockchain.ChainTips {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETCHAINTIPS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	chainTips := blockchain.ChainTips{}

	res, ok := rpcResult.Data.Result.([]interface{})
	if ok {
		for i := 0; i < len(res); i++ {
			var tip blockchain.Tip
			m, ok := res[i].(map[string]interface{})
			if ok {
				tip.Height = m["height"].(float64)
				tip.Hash = m["hash"].(string)
				tip.Branchlen = m["branchlen"].(float64)
				tip.Status = m["status"].(string)
			}
			chainTips.Tips = append(chainTips.Tips, tip)
		}
	}

	return chainTips
}

//获取区块链的交易状态
func (bc btcSer) GetChainTxStats() blockchain.ChainTxStats {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETCHAINTXSTATS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	chainTxStats := blockchain.ChainTxStats{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		chainTxStats.Window_final_block_hash = res["window_final_block_hash"].(string)
		chainTxStats.Window_final_block_height = res["window_final_block_height"].(float64)
		chainTxStats.Window_block_count = res["window_block_count"].(float64)
		chainTxStats.Window_tx_count = res["window_tx_count"].(float64)
		chainTxStats.Window_interval = res["window_interval"].(float64)
		chainTxStats.Txrate = res["txrate"].(float64)
		chainTxStats.Time = res["time"].(float64)
		chainTxStats.Txcount = res["txcount"].(float64)
	}

	return chainTxStats
}

//获取当前挖矿难度,成功返回挖矿难度,否则返回-1
func (bc btcSer) GetDifficulty() float64 {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETDIFFICULTY, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -0
}

//获取回收内存信息
func (bc btcSer) GetMempoolInfo() blockchain.MempoolInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETMEMPOOLINFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	mempoolInfo := blockchain.MempoolInfo{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		mempoolInfo.Loaded = res["loaded"].(bool)
		mempoolInfo.Size = res["size"].(float64)
		mempoolInfo.Bytes = res["bytes"].(float64)
		mempoolInfo.Usage = res["usage"].(float64)
		mempoolInfo.Maxmempool = res["maxmempool"].(float64)
		mempoolInfo.Mempoolminfee = res["mempoolminfee"].(float64)
		mempoolInfo.Minrelaytxfee = res["minrelaytxfee"].(float64)
	}

	return mempoolInfo
}

//author :陈浩亮  time ：2020/12/24
//分析psbt
func (bc btcSer) AnalyzePsbt(psbt string) rawTransactions.AnalyzePsbt {
	paramsSlice := []interface{}{psbt}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.ANALYZEPSBT , paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL,Rpc. RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	analyzePsbt := rawTransactions.AnalyzePsbt{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		analyzePsbt.Estimated_vsize = res["estimated_vsize"].(float64)
		analyzePsbt.Estimated_feerate = res["estimated_feerate"].(float64)
		analyzePsbt.Fee = res["fee"].(float64)
		analyzePsbt.Next = res["next"].(string)
		analyzePsbt.Error = res["error"].(string)
	}

	return analyzePsbt
}

//合并结合pspt
func (bc btcSer) CombinePsbt(txs string) string {
	paramsSlice := []interface{}{txs}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.COMBINEPSBT, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//合并原始交易
func (bc btcSer) CombineRawTransaction(txs string) string {
	paramsSlice := []interface{}{txs}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.COMBINERAWTRANSACTION, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//提取最终的psbt
func (bc btcSer) FinalizePsbt(psbting string) rawTransactions.FinalizePsbt {
	paramsSlice := []interface{}{psbting}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.FINALIZEPSBT , paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult :=Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	finalizePsbt := rawTransactions.FinalizePsbt{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		finalizePsbt.Psbt = res["psbt"].(string)
		finalizePsbt.Hex = res["hex"].(string)
		finalizePsbt.Complete = res["complete"].(bool)
	}

	return finalizePsbt
}

//转换为psbt
func (bc btcSer) ConverttoPsbt(rawTransaction string) string {
	paramsSlice := []interface{}{rawTransaction}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.CONVERTTOPSBT, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult :=Rpc. DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//找原始交易信息
func (bc btcSer) FundRawTransaction(rawTransaction string) rawTransactions.FundRawTransaction {
	paramsSlice := []interface{}{rawTransaction}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.FUNDRAWTRANSACTION , paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	fundRawTransaction := rawTransactions.FundRawTransaction{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		fundRawTransaction.Result = res["result"].(string)
		fundRawTransaction.Fee = res["fee"].(float64)
		fundRawTransaction.Changepos = res["changepos"].(float64)
	}

	return fundRawTransaction
}

//发送原始交易信息
func (bc btcSer) SendRawTransaction(rawTransaction string) string {
	paramsSlice := []interface{}{rawTransaction}
	//RPC通信标椎格JSON式数据
	rpcNormJson :=Rpc. PrepareJSON(utils.SENDRAWTRANSACTION, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//用私钥签名交易
func (bc btcSer) SignRawTransactionWithKey(pri string) rawTransactions.SignRawTransactionWithKey {
	paramsSlice := []interface{}{pri}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.SIGNRAWTRANSACTIONWITHKEY, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult :=Rpc. DoPost(utils.RPCURL,Rpc. RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	signRawTransactionWithKey := rawTransactions.SignRawTransactionWithKey{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		signRawTransactionWithKey.Hex = resBytes["hex"].(string)
		signRawTransactionWithKey.Complete = resBytes["complete"].(bool)
		signRawTransactionWithKey.Errors_, ok = resBytes["errors"].([]interface{})
		if ok {
			for i := 0; i < len(signRawTransactionWithKey.Errors_); i++ {
				mapValue, ok := signRawTransactionWithKey.Errors_[i].(map[string]interface{})
				if ok {
					var err rawTransactions.Error
					err.Txid = mapValue["txid"].(string)
					err.Vout = mapValue["vout"].(float64)
					err.ScriptSig = mapValue["scriptSig"].(string)
					err.Sequence = mapValue["sequence"].(float64)
					err.Error = mapValue["error"].(string)
					signRawTransactionWithKey.Errors = append(signRawTransactionWithKey.Errors, err)
				}
			}
		}


	}

	return signRawTransactionWithKey
}

//测试连接池是否接受链接
func (bc btcSer) TestMempoolAccept() rawTransactions.TestMempoolAccept {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.TESTMEMPOOLACCEPT, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL,Rpc. RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	testMempoolAccept := rawTransactions.TestMempoolAccept{}
	_, ok := rpcResult.Data.Result.(map[string]interface{})

	if ok {

		for i := 0; i < len(testMempoolAccept.Test_); i++ {
			mapValue, ok := testMempoolAccept.Test_[i].(map[string]interface{})
			if ok {
				var testing rawTransactions.Testing
				testing.Txid = mapValue["txid"].(string)
				testing.Allowed = mapValue["allowed"].(bool)
				testing.Reject_reason = mapValue["reject_reason"].(string)
				testMempoolAccept.Test = append(testMempoolAccept.Test, testing)
			}
		}
	}

	return testMempoolAccept
}

//创建多重签名需求
func (bc btcSer) CreateMultisig(nrequired  float64,pubkey string) util.CreateMultiSig {
	paramsSlice := []interface{}{nrequired,pubkey}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.CREATEMULTISIG , paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	createMultiSig := util.CreateMultiSig{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		createMultiSig.Address = res["address"].(string)
		createMultiSig.RedeemScript = res["redeemScript"].(string)
		createMultiSig.Descriptor = res["descriptor"].(string)
	}

	return createMultiSig
}

//地址起源
func (bc btcSer) DeriveAddresses(descriptor  string) string {
	paramsSlice := []interface{}{descriptor}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.DERIVEADDRESSES, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//估算费用
func (bc btcSer) EstimateSmartFee(conf_target float64) util.EstimateSmartFee {
	paramsSlice := []interface{}{conf_target}
	//RPC通信标椎格JSON式数据
	rpcNormJson :=Rpc. PrepareJSON(utils.ESTIMATESMARTFEE , paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL,Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	estimateSmartFee := util.EstimateSmartFee{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		estimateSmartFee.Feerate = res["feerate"].(float64)
		estimateSmartFee.Error = res["error"].(string)
		estimateSmartFee.Block = res["block"].(float64)
	}

	return estimateSmartFee
}

//获取描述符信息
func (bc btcSer) GetDesCriptorInfo(descriptor string) util.DesCriptorInfo {
	paramsSlice := []interface{}{descriptor}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETDESCRIPTORINFO , paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	desCriptorInfo := util.DesCriptorInfo{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		desCriptorInfo.Descriptor = res["descriptor"].(string)
		desCriptorInfo.Checksum = res["checksum"].(string)
		desCriptorInfo.Isrange = res["isrange"].(bool)
		desCriptorInfo.Issolvable = res["issolvable"].(bool)
		desCriptorInfo.Hasprivatekeys = res["hasprivatekeys"].(bool)
	}

	return desCriptorInfo
}

//用私钥对交易进行签名
func (bc btcSer) SignMessageWithprivKey(privkey  string) string {
	paramsSlice := []interface{}{privkey}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.SIGNMESSAGEWITHPRIVKEY, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}
	return ""
}

//验证地址信息
func (bc btcSer) ValidateAddress(address string) util.ValidateAddressInfo {
	paramsSlice := []interface{}{address}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.VALIDATEADDRESS , paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	//反序列化操作
	validateAddressInfo := util.ValidateAddressInfo{}

	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		validateAddressInfo.Isvalid = res["isvalid"].(bool)
		validateAddressInfo.Address = res["address"].(string)
		validateAddressInfo.ScriptPubKey = res["scriptPubKey"].(string)
		validateAddressInfo.Isscript = res["isscript"].(bool)
		validateAddressInfo.Iswitness = res["iswitness"].(bool)
		validateAddressInfo.Witness_version = res["witness_version"].(float64)
		validateAddressInfo.Witness_program = res["witness_program"].(string)
	}

	return validateAddressInfo
}

//----------------------------begin:wallet---------------------------------//
/*
 *Author： Xia
 *周日 12月 27 22:06 2020
 */
//返回关于重新扫描
func (bc btcSer) AbortRescan() bool {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.ABORTRESCAN, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(bool)
	if ok {
		return res
	}

	return false
}

//添加多重账户地址
func (bc btcSer) AddMultisigAddress(nrequried int64, keys []string) wallet.AddMultisigAddress {
	paramsSlice := []interface{}{nrequried, keys}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.ADDMULTISIGADDRESS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	multisigAddress := wallet.AddMultisigAddress{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		multisigAddress.Address = res["address"].(string)
		multisigAddress.RedeemScript = res["redemscript"].(string)
		multisigAddress.Descriptor = res["descriptor"].(string)
	}
	return multisigAddress
}


//查询撞的费用
func (bc btcSer) BumpFee(txId string) wallet.BumpFee {
	paramsSlice := []interface{}{txId}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.BUMPFEE, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	bumpfeeIfo := wallet.BumpFee{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		bumpfeeIfo.Psbt = res["psbt"].(string)
		bumpfeeIfo.TxId = res["redemscript"].(string)
		bumpfeeIfo.Origfee = res["origfee"].(int64)
		bumpfeeIfo.Fee = res["origfee"].(int64)
		bumpfeeIfo.Errors = res["errorx"].([]string)
		bumpfeeIfo.Str = res["str"].([]string)
	}
	return bumpfeeIfo
}

//=创建钱包
func (bc btcSer) CreateWallet(wallet_name string, passphrase string) wallet.CreateWallet {
	paramsSlice := []interface{}{wallet_name, passphrase}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.CREATEWALLET, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	createWallet := wallet.CreateWallet{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		createWallet.Name = res["name"].(string)
		createWallet.Warning = res["warning"].(string)
	}
	return createWallet
}

//转储私钥
func (bc btcSer) DumpPrivkey(adress string) string {
	paramsSlice := []interface{}{adress}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.DUMPPRIVKEY, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res

	}
	return ""
}

//转储钱包
func (bc btcSer) DumpWallet(filename string) string {
	paramsSlice := []interface{}{filename}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.DUMPWALLET, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res

	}
	return ""
}

//加密钱包
func (bc btcSer) EncyptWallet(passphrase string) string {
	paramsSlice := []interface{}{passphrase}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.ENCRYPTWALLET, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res

	}
	return ""
}



//获取地址信息
func (bc btcSer) GetAddressInfo(address string) wallet.AddressInfo {
	paramsSlice := []interface{}{address}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETADDRESSINFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	getAddressInfo := wallet.AddressInfo{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		getAddressInfo.Address = res["address"].(string)
		getAddressInfo.ScriptPubKey = res["scriptPubKey"].(string)
		getAddressInfo.Ismine = res["ismine"].(bool)
		getAddressInfo.Solvable = res["solvable"].(bool)
		getAddressInfo.Desc = res["desc"].(string)
		getAddressInfo.Iswatchonly = res["iswatchonly"].(bool)
		getAddressInfo.Isscript = res["isscript"].(bool)
		getAddressInfo.Iswitness = res["iswitness"].(bool)
		//getAddressInfo.Hex = res["hex"].(nil)
		getAddressInfo.Pubkey = res["pubkey"].(string)
		getAddressInfo.Ischange = res["ischange"].(bool)
		getAddressInfo.Timestamp = res["timestamp"].(float64)
		//getAddressInfo.Hdkeypath = res["hdkeypath"].(string)
		getAddressInfo.Hdseedid = res["hdseedid"].(string)
		getAddressInfo.Hdmasterfingerprint = res["hdmasterfingerprint"].(string)
		getAddressInfo.Labels_ = res["labels"].([]interface{})

		getAddressInfo.Embedded_, ok = res["Embedded"].(map[string]interface{})
		if ok {
			getAddressInfo.Embedded.Isscript = res["isscript"].(bool)
			getAddressInfo.Embedded.Iswitness = res["iswitness"].(bool)
			getAddressInfo.Embedded.Witness_version = res["witness_version"].(float64)
			getAddressInfo.Embedded.Witness_program = res["witness_program"].(string)
			getAddressInfo.Embedded.Pubkey = res["pubkey"].(string)
			getAddressInfo.Embedded.Address = res["address"].(string)
			getAddressInfo.Embedded.ScriptPubKey = res["scriptpubkey"].(string)
		}

	}
	return getAddressInfo
}

//返回这个钱包收到的比特币总数
func (bc btcSer) GetBalance() float64 {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.ABORTRESCAN, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -1
}

//返回一个BTC中所有余额的对象
func (bc btcSer) GetBalances() wallet.Balances {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.ABORTRESCAN, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	getBalances := wallet.Balances{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		getBalances.Mine_, ok = res["Mine"].(map[float64]interface{})
		if ok {
			getBalances.Mine.Trusted = res["trusted"].(float64)
			getBalances.Mine.Untrusted_pending = res["untrusted_pending"].(float64)
			getBalances.Mine.Immature = res["immature"].(float64)
			getBalances.Mine.Used = res["used"].(float64)
		}
		getBalances.Watchonly_, ok = res["Watchonly"].(map[float64]interface{})
		if ok {
			getBalances.Mine.Trusted = res["trusted"].(float64)
			getBalances.Mine.Untrusted_pending = res["untrusted_pending"].(float64)
			getBalances.Mine.Immature = res["immature"].(float64)
		}
	}

	return getBalances
}

//返回新地址
func (bc btcSer) GetNewAddress() string {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETNEWADDRESS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}

	return ""
}

//获取原始更改地址
func (bc btcSer) GetRawChangeAddress() string {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETRAWCHANGEADDRESS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}

	return ""
}

//返回给定地址接收到的总金额
func (bc btcSer) GetReceivedByAddress(address string) float64 {
	paramsSlice := []interface{}{address}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETRECEIVEDBYADDRESS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -1
}
//=返回收到标签
func (bc btcSer) GetReceivedByLabel(label string) float64 {
	paramsSlice := []interface{}{label}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETRECEIVEDBYLABEL, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -1
}
//返回有未经证实的平衡的比特币
func (bc btcSer) GetUnconfirmedBalance() float64 {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETRECEIVEDBYLABEL, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(float64)
	if ok {
		return res
	}

	return -1
}
//返回钱包信息
func (bc btcSer) GetWalletInfo()wallet.WalletInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETWALLETINFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))
	getWallentInfo := wallet.WalletInfo{}
	res, ok := rpcResult.Data.Result.(map[string]interface{})
	if ok {
		getWallentInfo.Walletname = res["walletname"].(string)
		getWallentInfo.Walletversion = res["walletversion"].(float64)
		getWallentInfo.Balance = res["balance"].(float64)
		getWallentInfo.Unconfirmed_balance = res["unconfirmed_balance"].(float64)
		getWallentInfo.Immature_balance = res["immature_balance"].(float64)
		getWallentInfo.Txcount = res["txcount"].(float64)
		getWallentInfo.Keypoololdest = res["keypoololdest"].(float64)
		getWallentInfo.Keypoolsize = res["keypoolsize"].(float64)
		getWallentInfo.Hdseedid = res["hdseedid"].(string)
		getWallentInfo.Keypoolsize_hd_internal = res["keypoolsize_hd_internal"].(float64)
		getWallentInfo.Paytxfee = res["paytxfee"].(float64)
		getWallentInfo.Private_keys_enabled = res["private_keys_enabled"].(bool)
		getWallentInfo.Avoid_reuse = res["avoid_reuse"].(bool)
		getWallentInfo.Scanning = res["scanning"].(bool)
	}

	return getWallentInfo
}
//=返回一个重要的地址
func (bc btcSer) ImportAddress(address string) string {
	paramsSlice := []interface{}{address}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.IMPORTADDRESS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}

	return ""
}
//返回标签列表
func (bc btcSer) ListLabels() string {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.LISTLABELS, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	res, ok := rpcResult.Data.Result.(string)
	if ok {
		return res
	}

	return ""
}
//------------------------------end--------------------------------//