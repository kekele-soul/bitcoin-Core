package control
/**
   *Author： zck
*周二 12月 22 20:50 2020
*/

//获得内存信息
type MemoryInfo struct {
	Used        float64
	Free        float64
	Total       float64
	Locked      float64
	Chunks_used float64
	Chunks_free float64
}
//得到rpc信息
type RpcInfo struct {
	Method   string
	Duration float64
	Logpath  string
}
//获取并设置日志记录配置
type Logging struct {
	Net         bool
	Tor         bool
	Mempool     bool
	Http        bool
	Bench       bool
	Zmq         bool
	Walletdb    bool
	Rpc         bool
	Estimatefee bool
	Addrman     bool
	Selectcoins bool
	Reindex     bool
	Cmpctblock  bool
	Rand        bool
	Prune       bool
	Proxy       bool
	Mempoolrej  bool
	Libevent    bool
	Coindb      bool
	Qt          bool
	Leveldb     bool
	Validation  bool
}
