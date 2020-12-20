//此处写RPC通信方法

package bitcoinServices

import (
	"bitcoin-Core/utils/Rpc"
	"bitcoin-Core/utils"
	"bitcoin-Core/models/blockchain"
	"strings"
)

//获取RPCURL链接节点的最高区块Hash
func GetBestBlockHahs() string {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBESTBLOCKHASH, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	return rpcResult.Data.Result.(string)
}

//根据区块Hash获取区块信息
func GetBlockInfo(blockHash string) blockchain.BlockInfo { //
	paramsSlice := []interface{}{blockHash}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCK, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	blockInfo := blockchain.BlockInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if !ok {
		return blockInfo
	}

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
	return blockInfo
}

//获取区块链信息
func GetBlockChainInfo() blockchain.BlockChainInfo {
	paramsSlice := []interface{}{}
	//RPC通信标椎格JSON式数据
	rpcNormJson := Rpc.PrepareJSON(utils.GETBLOCKCHAININFO, paramsSlice)

	//bitcoin Core 响应的结果
	rpcResult := Rpc.DoPost(utils.RPCURL, Rpc.RequestHeaders(), strings.NewReader(rpcNormJson))

	blockChainInfo := blockchain.BlockChainInfo{}
	resBytes, ok := rpcResult.Data.Result.(map[string]interface{})
	if !ok {
		return blockChainInfo
	}

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
	return blockChainInfo
}