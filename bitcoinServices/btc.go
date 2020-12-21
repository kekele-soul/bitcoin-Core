package bitcoinServices

import (
	"bitcoin-Core/models/blockchain"
)

type BlockChahin interface {
	//获取最高区块Hash
	GetBestBlockHahs() string

	//根据区块Hash获取区块信息
	GetBlockInfoByHash(blockHash string) blockchain.BlockInfo

	//根据区块hash获取区块Info
	GetBlockByHeight(height int) blockchain.BlockInfo

	//获取区块链信息
	GetBlockChainInfo() blockchain.BlockChainInfo

	//获取网络信息
	GetNetWorkInfo() blockchain.NetWorkInfo

	//获取区块总数
	GetBlockCount() float64

	//根据区块高度获取区块hash
	GetBlockHash(height int) string

	//根据区块Hash获取区块头信息
	GetBlockHeaderInfoByHash(hash string) blockchain.BlockHeaderInfo

	//根据区块Height获取区块头信息
	GetBlockHeaderInfoByHeight(height float64) blockchain.BlockHeaderInfo

	//根据区块高度获取区块状态
	GetBlockStatsInfoByHeight(height float64) blockchain.BlockStats

	//根据区块Hash获取区块状态
	GetBlockStatsInfoByHash(hash string) blockchain.BlockStats

	//获取区块链Tips
	GetChainTips() blockchain.ChainTips

	//获取区块链的交易状态
	GetChainTxStats() blockchain.ChainTxStats

	//获取当前挖矿难度
	GetDifficulty() float64

	//根据TxId获取到内存池的祖先
	GetMempoolAncestors(txId string) blockchain.MempoolAncestorsInfo

	//根据TxId获取到内存池的后代
	GetMempoolDescendants(txId string) blockchain.MempoolDescendantsInfo

	//根据TxId获取内存池数据
	GetMempoolEntry(txId string) blockchain.MempoolEntryInfo

	//获取回收内存信息
	GetMempoolInfo() blockchain.MempoolInfo
}
