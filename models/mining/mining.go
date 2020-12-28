package mining

/**
   *Author： zck
*周二 12月 22 20:50 2020
*/

//得到挖掘信息
type MiningInfo struct {
	Blocks        float64
	Difficulty    float64
	Networkhashps float64
	Pooledtx      string
	Warnings      string
}

type PrioritiseTransaction struct {
}

type SubmitBlock struct {
}

type SubmitHeader struct {
}

type NetWorkInfo struct {
}
