package Convert

/**
该包封装为将int64和string转化为byte字节切片工具包
 */

import (
	"bytes"
	"encoding/binary"
)

/**
	将int64转化为[]byte
 */
func IntToBytes(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	//大端位序排列：binary.BigEndian
	//小端位序排列：binary.LittleEndian
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

/**
	将string转化为[]byte
 */
func StringToBytes(st string) []byte {
	return []byte(st)
}
