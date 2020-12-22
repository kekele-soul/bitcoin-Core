/*
 *Author： Afei
 *周二 12月 22 15：09 2020
 */
package db

import (
	"bitcoin-Core/utils/Base"
	"bitcoin-Core/utils/Hash"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

//用户名和密码等数据操作

var Map_Name_Pwd map[string]string

//保存用户名和密码
func SaveUser(path string, name, pwd string) error {

	if name == "" {
		return errors.New("用户名不能为空!")
	}

	//打开文件，
	f, err := os.OpenFile(path, os.O_WRONLY, 0644) //传递文件路径
	if err != nil {                                //有错误
		return err
	}
	//计算偏移量
	n, err := f.Seek(0, 2)
	if err != nil {
		return err
	}

	//使用完毕，需要关闭文件
	defer f.Close()
	//初始化map
	err = Init(path)
	if err != nil {
		return err
	}

	//判断用户名是否存在
	if Map_Name_Pwd[name] != "" {
		str := fmt.Sprintf("注册失败,用户名:%v 已存在!", name)
		return errors.New(str)
	}
	//对密码进行hash
	pwdBytes := Hash.Sha256Hash([]byte(pwd))
	//对hash后的密码进行base64编码
	base64PwdBytes := Base.Base64Encode(pwdBytes)
	//对编码后的密码转换为string类型
	pwdStr := string(base64PwdBytes)
	a := name + "," + pwdStr + "\n"
	_, err = f.WriteAt([]byte(a), n)
	if err != nil {
		return err
	}
	//初始化map
	err = Init(path)
	if err != nil {
		return err
	}
	return nil
}

//查询用户
func Query(name string, pwd string) (bool, error) {
	err := Init()
	if err != nil {
		return false, err
	}
	mp := Map_Name_Pwd

	//判断用户是否存在
	pwdBase := mp[name]
	if pwdBase == "" {
		return false, errors.New("用户不存在")
	}

	//对用户的密码pwd进行hash和base编码
	pwdBytes := Hash.Sha256Hash([]byte(pwd))
	pwdBase64Bytes := Base.Base64Encode(pwdBytes)
	pwdBaseStr := string(pwdBase64Bytes)

	//比较密码是否相等并返回结果
	return pwdBase == pwdBaseStr, nil
}

//初始化map
func Init(path string) error {
	Map_Name_Pwd = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	//建立缓冲区，把文件内容放到缓冲区中
	buf := bufio.NewReader(f)
	for {
		//遇到\n结束读取
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			return err
		}

		filr := strings.Split(string(b), "\n")

		f1 := filr[0]
		br := strings.Split(f1, ",")

		for i := 0; i < len(br); i++ {
			if i == 1 {
				if Map_Name_Pwd[br[0]] == "" {
					Map_Name_Pwd[br[0]] = br[i]
				}
			}
		}

	}
	return nil
}