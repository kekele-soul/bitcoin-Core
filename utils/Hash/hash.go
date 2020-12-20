package Hash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
)

/**
 *MD5哈希应用
 */

/**
字节进行MD5哈希
*/
func Md5Hash(data []byte) []byte {
	Md5Hash := md5.New()
	Md5Hash.Write(data)
	return Md5Hash.Sum(nil)
}

/**
字符串进行MD5哈希
*/
func MD5HashString(data string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	psswordBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(psswordBytes)
}

/**
文件进行MD5哈希
*/
func MD5HashReader(reader io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}


/**
字节进行sha256哈希
*/
func Sha256Hash(data []byte) []byte {
	Sha256Hash := sha256.New()
	Sha256Hash.Write(data)
	return Sha256Hash.Sum(nil)
}

/**
字符串进行sha256哈希
*/
func Sha256HashString(data string) string {
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(data))
	Bytes := sha256Hash.Sum(nil)
	return hex.EncodeToString(Bytes)
}

/**
文件进行sha256哈希
*/
func SHA256HashReader(reader io.Reader) (string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	sh256Hash := sha256.New()
	sh256Hash.Write(bytes)
	return hex.EncodeToString(sh256Hash.Sum(nil)), nil
}
