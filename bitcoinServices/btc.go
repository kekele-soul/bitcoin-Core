package bitcoinServices

import (
	"bitcoin-Core/models/blockChain"
	"bitcoin-Core/models/control"
	"bitcoin-Core/models/generating"
	"bitcoin-Core/models/mining"
	"bitcoin-Core/models/network"
	"bitcoin-Core/models/rawTransactions"
	"bitcoin-Core/models/util"
	"bitcoin-Core/models/wallet"
)

type BlockChahin interface {
	//获取最高区块Hash
	GetBestBlockHash() string

	//根据区块Hash获取区块信息
	GetBlockInfoByHash(blockHash string) blockChain.BlockInfo

	//根据区块hash获取区块Info
	GetBlockInfoByHeight(height int) blockChain.BlockInfo

	//获取区块链信息
	GetBlockChainInfo() blockChain.BlockChainInfo

	//获取网络信息
	GetNetWorkInfo() blockChain.NetWorkInfo

	//获取区块总数
	GetBlockCount() float64

	//根据区块高度获取区块hash
	GetBlockHash(height int) string

	//根据区块Hash获取区块头信息
	GetBlockHeaderInfoByHash(hash string) blockChain.BlockHeaderInfo

	//根据区块Height获取区块头信息
	GetBlockHeaderInfoByHeight(height float64) blockChain.BlockHeaderInfo

	//根据区块高度获取区块状态
	GetBlockStatsInfoByHeight(height float64) blockChain.BlockStats

	//根据区块Hash获取区块状态
	GetBlockStatsInfoByHash(hash string) blockChain.BlockStats

	//获取区块链Tips
	GetChainTips() blockChain.ChainTips

	//获取区块链的交易状态
	GetChainTxStats() blockChain.ChainTxStats

	//获取当前挖矿难度
	GetDifficulty() float64

	//根据TxId获取到内存池的祖先
	GetMempoolAncestors(txId string) blockChain.MempoolAncestorsInfo

	//根据TxId获取到内存池的后代
	GetMempoolDescendants(txId string) blockChain.MempoolDescendantsInfo

	//根据TxId获取内存池数据
	GetMempoolEntry(txId string) blockChain.MempoolEntryInfo

	//获取回收内存信息
	GetMempoolInfo() blockChain.MempoolInfo
}

//一个人完成以下四个接口
//控制
type Control interface {
	//获得内存信息
	GetMemoryInfo() control.MemoryInfo

	//返回RPC服务器的详细信息
	GetRpcInfo() control.RpcInfo
	//查询所以命令
	Help() string

	//获取和设置日志记录配置。
	//在不带参数的情况下调用时，返回状态是否为调试日志的类别列表。
	LogGing() control.LogGingInfo

	//关闭Bitcoin Cone
	Stop()

	//Bitcoin Cone正常运行时间一秒为单位
	UpTime() float64
}

//生产
type Generating interface {
	//直接将块挖掘到指定的地址(在RPC调用返回之前)
	GenerateToAddress(nblocks int64, address string) generating.Generatetoaddress
	//立即将块挖掘到指定的描述符(在RPC调用返回之前)
	GenerateToDescriptor(num_blocks int64, descriptor string) generating.Generatetodescriptor
}

//矿工
type Mining interface {
	//得到挖掘信息
	GetMiningInfo() mining.MiningInfo

	//获取每秒的网络散列  PS:per second
	GetNetWorkHashPS() float64

	GetNetWorkInfo() mining.NetWorkInfo
	//以更高的优先级(或更低的优先级)接受事务到被挖掘的块中
	PrioritiseTransaction(txId string) mining.PrioritiseTransaction
	//试图向网络提交新的块。
	Submitblock(hexdata string, dummy string) mining.SubmitBlock
	//解码给定的hexdata作为头和提交它作为一个候选链提示，如果有效。头无效时抛出。
	Submitheader(hexdata string) mining.SubmitHeader
}

//网络
type NetWork interface {
	//
	AddNode(node string, command string) network.AddNode

	//明确禁止
	ClearBanned() string

	//断距nect节点
	DisConnectNode(address string, nodeid int64) network.DisconnectNode

	//获取添加节点信息
	GetAddednodeInfo() []network.AddedNodeInfo

	//获得连接数
	GetConnecTionCount() float64

	//获得净租金
	GetnNettoTals() network.NetTotals

	//获取网络信息
	GetNetWorkInfo() network.NetWorkInfo

	//得到节点地址
	GetNodeAddResses() network.NodeAddresses

	//以json对象数组的形式返回有关每个连接网络节点的数据。
	GetPeerInfo() string

	//List all manually banned IPs/Subnets.
	//列出所有手动禁用的ip /子网。N
	ListBanned() network.ListBanned

	Ping() string

	SetBan() string

	//禁用/启用所有p2p网络活动。
	SetNetWorkActive() bool
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
	CreateMultisig(nrequired float64, pubkey string) util.CreateMultiSig

	//地址起源
	DeriveAddresses(descriptor string)

	//估算费用
	EstimateSmartFee(conf_target float64) util.EstimateSmartFee

	//获取描述符信息
	GetDesCriptorInfo(descriptor string) util.DesCriptorInfo

	//用私钥对交易进行签名
	SignMessageWithprivKey(privkey string)

	//验证地址信息
	ValidateAddress(address string) util.ValidateAddressInfo
}

//夏易阳、何欣萍两人合作完成下面一个接口
//钱包
type Wallet interface {
	//放弃交易
	AbandonTransaction(txId string) wallet.AbandonTransactionInfo

	//返回关于重新扫描
	AbortRescan() wallet.AbortRescan

	//添加多重账户地址
	AddMultisigAddress(nrequried int64, keys []string) wallet.AddMultisigAddress

	//转至备份钱包
	BackUpWallet(destination string) wallet.UpWallet

	//查询撞的费用
	BumpFee(txId string) wallet.BumpFee

	//创建钱包
	CreateWallet(wallet_name string, passphrase string) wallet.CreateWallet

	//转储私钥
	DumpPrivkey(adress string) wallet.DumpPrivkey

	//转储钱包
	DumpWallet(filename string)

	//加密钱包
	EncyptWallet(passphrase string)

	//通过标签获取地址
	//GetAddressesByLabel(label string) entity.Label

	//获取地址信息
	GetAddressInfo(address string) wallet.AddressInfo

	//返回这个钱包收到的比特币总数
	GetBalance()

	//返回一个BTC中所有余额的对象。
	GetBalances()

	//返回新地址
	GetNewAddress()

	//获取原始更改地址
	GetRawChangeAddress()

	//返回给定地址接收到的总金额。
	GetReceivedByAddress()

	//返回收到标签
	GetReceivedByLabel()

	//返回有未经证实的平衡的比特币
	GetUnconfirmedBalance()

	//返回钱包信息
	GetWalletInfo() wallet.WalletInfo

	//返回一个重要的地址
	ImportAddress()

	//重要的
	ImportMulti()

	//返回重要的私钥
	ImportPrivkey()

	//返回修理的资金
	ImportPrunedFunds()

	//返回重要的公钥
	ImportPurvkey()

	//返回重要的钱包
	ImportWallet()

	//填充钥匙池
	KeyPoolRefill()

	//返回地址列表分组
	ListAddressGroupings()

	//返回标签列表
	ListLabels()

	/*
	 *将金额发送到指定地址.如果钱包被加密，需要用WalletPassPhrase()调用来设置钱包密码.
	 *address:接受BTC的地址.必填
	 *amount:发送的BTC金额,如0.1.必填
	 *comment:用于存储交易用途的注释.可选
	 *commit_to:用于存储个人或组织名称的注释,就放在你的钱包里.可选
	 *subtractfeefromamount:交易费用是否从汇款金额中扣除.可选,默认为false
	 *replaceable:通过BIP 125允许该交易被更高费用的交易取代.可选,默认为"wallet default"
	 *conf_target:确认目标.可选,默认为"wallet default"
	 *estimate_mode:可选,默认为UNSET.必须为一下值:"UNSET"、"ECONOMICAL"、"CONSERVATIVE".
	 *avoid_reuse:避免从肮脏的地址消费.可选,默认为true
	 *result:返回交易ID
	 */
	SendToAddress(address string, amount float64, comment, commit_to string, subtractfeefromamount, replaceable bool, conf_target float64, estimate_mode string, avoid_reuse bool) string

	/*
		 *设置或生成一个新的HD钱包种子.
		 *非高清钱包不会升级为高清钱包. 已经存在的钱包HD将会有一个新的HD种子集，这样新添加到密钥池的密钥将会从这个新的种子衍生出来。
		 *请注意，在设置了HD钱包种子后，您需要重新备份您的钱包.
		 *如果钱包被加密，需要用WalletPassPhrase()调用来设置钱包密码.
		 *newKeyPool:默认填true.是否从密钥池中刷新旧的未使用地址(包括更改地址)并重新生成它.
			如果为真，下一个来自GetNewAddress()的地址和来自GetRawChangeAddress()的ChangeAddress()将来自这个新种子.
			如果为false，从现有地址(包括更改地址，如果钱包已经有HD链拆分启用)密钥池将一直使用到耗尽为止.
		 *seed:新的HD种子使用的WIF私钥。
		 *result:null
	*/
	SetHdSeed(newKeyPool bool, seed string)

	/*
	 *设置与给定地址相关联的标签.
	 *address:设置标签的地址
	 *label:标签
	 *result:null
	 */
	SetLabel(address string, label string)

	/*
	 *为这个钱包设置每kB的交易费用。
	 *amount:数量.以BTC/kB为单位的交易费用
	 *result:如果成功则返回true,否则返回false
	 */
	SetTxFee(amount float64) bool

	/*
	 *更改钱包的给定钱包标志的状态。
	 *flag:要更改的标志的名称.
	 *value:是否新规定.默认为填写true
	 *result:
	 */
	SetWalletFlag(flag string, value bool) wallet.WalletFlagInfo

	/*
	 *使用地址的私钥签署消息.如果钱包被加密，需要用WalletPassPhrase()调用来设置钱包密码。
	 *address:私钥对应的比特币地址。
	 *message:要创建签名的消息。
	 *result:以base 64编码的消息签名
	 */
	SignMessage(address string, message string) string

	/*
		 *用密钥对原始交易进行签名.
		 *hexString:交易十六进制字符串.
		 *privkeys:用base58编码的用于签名的私钥.
		 *prevtxs:先前的依赖交易输出.可选, 默认给空切片
		 *sighashtype:签名哈希类型.可选, 默认为ALL
		必须是一下类型:"ALL"、"NONE"、"SINGLE"、"ALL|ANYONECANPAY"、"NONE|ANYONECANPAY"、"SINGLE|ANYONECANPAY"
	*/
	SignRawTransactionWithKey(hexString string, privkeys []string, prevtxs []wallet.PrevTx, sighashtype string) wallet.SignResult

	/*
	 *卸载请求端点引用的钱包，否则卸载参数中指定的钱包.
	 *wallet_name: 可选, 默认为:来自RPC请求的钱包名称
	 */
	//Examples: UnLoadWallet(wallet_name)-->
	UnLoadWallet(wallet_name ...string)

	/*
	 *实现创建者和更新者角色.
	 *inputs:必须要的参数.
	 *outputs:必须要的参数,.第一个为entity.Address类型的结构体,第二个开始为entity.Data类型的结构体
	 *locktime:可选可选的参数.默认填零
	 *obj:可选的参数.默认为nil
	 *bip32derivs:可选的参数.默认为true
	 */
	//Examples: WalletCreateFundedPsbt([...], [...], 0, nilObj, true)
	WalletCreateFundedPsbt(inputs []wallet.Input, outputs []interface{}, locktime int, obj wallet.Obj, bip32derivs bool)

	//清除钱包锁.
	//Examples:WalletPassPhrase()-->SendToAddress()-->WalletLock()
	WalletLock()

	//为执行一个交易设置timeOut时间的密码,timeOut:以秒为单位
	WalletPassPhrase(PassPhrase string, timeOut int64)

	//更改钱包密码
	WalletPassPhraseChange(oldPassPhrase, newPassPhrase string)

	WalletProcessPsbt(psbt string) wallet.Processps
}

//消息队列
type Zmq interface {
}
