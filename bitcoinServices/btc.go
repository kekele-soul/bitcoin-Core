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

//一个人完成以下四个接口
//控制
type Control interface {

}

//生产
type Generating interface {

}

//矿工
type Mining interface {

}

//网络
type Network interface {

}


//一个人完成以下两个接口
//原始交易
type Rawtransactions interface {

}

//工具
type Util interface {

}


//两人合作完成下面一个接口
//钱包
type Wallet interface {

}


//消息队列
type Zmq interface {

}