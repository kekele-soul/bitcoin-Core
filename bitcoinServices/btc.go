package bitcoinServices

import (
	"bitcoin-Core/models/blockchain"
	"bitcoin-Core/models/control"
	"bitcoin-Core/models/generating"
	"bitcoin-Core/models/mining"
	"bitcoin-Core/models/network"
	"bitcoin-Core/models/rawTransactions"
	"bitcoin-Core/models/util"
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
	Getmemoryinfo() control.MemoryInfo
	//得到rpc信息
	Getrpcinfo() control.RpcInfo
	//查询所以命令
	Help() string
	//获取并设置日志记录配置
	Logging() control.Logging
	//要求体面地关闭比特币核心，
	Stop() string
	//返回服务器的总正常运行时间
	Uptime() float64
}

//生产
type Generating interface {
	//直接将块挖掘到指定的地址(在RPC调用返回之前)
	Generatetoaddress(nblocks int64, address string) generating.Generatetoaddress
	//立即将块挖掘到指定的描述符(在RPC调用返回之前)
	Generatetodescriptor(num_blocks int64, descriptor string) generating.Generatetodescriptor
}

//矿工
type Mining interface {
	//得到挖掘信息
	Getmininginfo() mining.MiningInfo
	//得到净工作哈希值ps
	Getnetworkhashps() float64

	Getnetworkinfo() mining.NetWorkInfo
	//以更高的优先级(或更低的优先级)接受事务到被挖掘的块中
	Prioritisetransaction(txId string) mining.Prioritisetransaction
	//试图向网络提交新的块。
	Submitblock(hexdata string, dummy string) mining.Submitblock
	//解码给定的hexdata作为头和提交它作为一个候选链提示，如果有效。头无效时抛出。
	Submitheader(hexdata string) mining.Submitheader
}

//网络
type NetWork interface {
	//
	AddNode(node string, command string) network.Addnode

	//明确禁止
	ClearBanned() string


	//断距nect节点
	DisConnectNode(address string, nodeid int64) network.Disconnectnode

	//获取添加节点信息
	GetAddednodeInfo() []network.AddedNodeInfo
	//获得连接数
	GetConnecTionCount() float64
	//获得净租金
	GetnNettoTals() network.Nettotals
	//获取网络信息
	GetNetWorkInfo() network.NetWorkInfo
	//得到节点地址
	GetNodeAddResses() network.NodeAddresses
	//以json对象数组的形式返回有关每个连接网络节点的数据。
	GetPeerInfo() string
	//List all manually banned IPs/Subnets.
	//列出所有手动禁用的ip /子网。N
	listBanned() network.Listbanned
	Ping() string
	SetBan() string
	//禁用/启用所有p2p网络活动。
	SetNetWorkacTive() bool
}


//author : 陈浩亮  time ：2020/12/24
//一个人完成以下两个接口
//原始交易
type Rawtransactions interface {
	//分析psbt
	AnalyzePsbt(psbt string) rawTransactions.AnalyzePsbt

	//合并结合pspt
	CombinePsbt(txs string)

	//合并结合原始交易
	CombineRawTransaction(txs string)

	//提取最终的psbt
	FinalizePsbt(psbting string) rawTransactions.FinalizePsbt

	//转换为psbt
	ConverttoPsbt(rawTransaction string)

	//找原始交易
	FundRawTransaction(rawTransaction string) rawTransactions.FundRawTransaction

	//发送原始交易信息
	SendRawTransaction(rawTransaction string)

	//用私钥签名交易
	SignRawTransactionWithKey(pri string) rawTransactions.SignRawTransactionWithKey

	//测试连接池是否接受链接
	TestMempoolAccept() rawTransactions.TestMempoolAccept
}

//工具
type Util interface {
	//创建多重签名需求
	CreateMultisig(nrequired  float64,pubkey string) util.CreateMultiSig

	//地址起源
	DeriveAddresses(descriptor  string)

	//估算费用
	EstimateSmartFee(conf_target float64) util.EstimateSmartFee

	//获取描述符信息
	GetDesCriptorInfo(descriptor string) util.DesCriptorInfo

	//用私钥对交易进行签名
	SignMessageWithprivKey(privkey string)

	//验证地址信息
	ValidateAddress(address string) util.ValidateAddressInfo

}


//两人合作完成下面一个接口
//钱包
type Wallet interface {

}


//消息队列
type Zmq interface {

}