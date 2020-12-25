package bitcoinServices

import (
	"bitcoin-Core/models/blockchain"
	"bitcoin/entity"
)

type BlockChahin interface {
	//获取最高区块Hash
	GetBestBlockHash() string

	//根据区块Hash获取区块信息
	GetBlockInfoByHash(blockHash string) blockchain.BlockInfo

	//根据区块hash获取区块Info
	GetBlockInfoByHeight(height int) blockchain.BlockInfo

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

//author : 陈浩亮  time ：2020/12/22
//一个人完成以下两个接口
//原始交易
type Rawtransactions interface {
	//分析psbt
	Analyzepsbt(psbt string)

	//合并结合pspt
	Combinepsbt() entity.Combinepsbt

	//合并结合原始交易
	Combinerawtransaction() entity.Combinerawtransaction

	//创建pspt
	Createpsbt() entity.Createpsbt

	//创建一个原始交易
	Createrawtransaction() entity.Createrawtransaction

	//找原始交易
	Fundrawtransaction(hexstring string)

	//发送原始交易信息
	Sendrawtransaction() entity.Sendrawtransaction

	//用私钥签名交易
	Signrawtransactionwithkey() entity.Signrawtransactionwithkey

	//测试连接池是否接受链接
	Testmempoolaccept() entity.Testmempoolaccept
}

//工具
type Util interface {
	//创建多重签名需求
	CreatemultisigNrequired() entity.CreatemultisigNrequired

	//地址起源
	Deriveaddresses() entity.Deriveaddresses

	//估算费用
	Estimate() entity.Estimate

	//获取描述符信息
	Getdescriptorinfo(descriptor string)

	//用私钥对交易进行签名
	Signmessagewithprivkey() entity.Signmessagewithprivkey

	//验证地址信息
	Validateaddress(address string)
}


//两人合作完成下面一个接口
//钱包
type Wallet interface {

}


//消息队列
type Zmq interface {

}