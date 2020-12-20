package Base

/**
该包为进行base64编码解码工具包
 */

import (
	"encoding/base64"
)


//base64编码
func Base64Encode(data []byte) []byte {
	encoding := base64.StdEncoding
	dst := make([]byte,encoding.EncodedLen(len(data)))
	encoding.Encode(dst, data)
	return dst

}

//base64解码
func Dase64Decode(data []byte) []byte {
	enCode:=base64.StdEncoding
	dst := make([]byte, enCode.DecodedLen(len(data)))
	n ,_:= enCode.Decode(dst,data)
	return dst[:n]

}


func Base64Str(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}

