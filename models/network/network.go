package network

/**
   *Author： zck
*周二 12月 22 20:50 2020
*/

type Addnode struct {
}

//断距nect节点
type Disconnectnode struct {
}
type Addresses struct {
	Address   string
	Connected string
}

//获取添加节点信息
type AddedNodeInfo struct {
	Addednode  string
	Connected  bool
	Addresses  []Addresses
	Addresses_ []interface{}
}

//获得净租金
type Nettotals struct {
	Totalbytesrecv float64
	Totalbytessent float64
	Timemillis     float64
	Uploadtarget   []interface{}
}
type Localservicesnames struct {
	WITNESS         string
	NETWORK_LIMITED string
}

//获取网络信息
type NetWorkInfo struct {
	Version            float64
	Subversion         string
	Protocolversion    float64
	Localservices      string
	Localservicesnames []interface{}
}

//得到节点地址
type NodeAddresses struct {
	Time     float64
	Services float64
	Address  string
	Post     float64
}

//List all manually banned IPs/Subnets.
type Listbanned struct {
	Address      string
	Banned_until float64
	Ban_created  float64
}

