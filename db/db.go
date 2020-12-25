/*
 *Author： Afei
 *周二 12月 22 15：09 2020
 */

//用户数据操作

package db

import (
	"bitcoin-Core/utils"
	"bitcoin-Core/utils/Base"
	"bitcoin-Core/utils/Hash"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var Map_Name_Pwd map[string]string
var Map_MailBox_Pwd map[string]string
var Map_Name_MailBox map[string]string
var Map_MailBox_Name map[string]string

//根据用户名登入
func QueryByName(name string, pwd string) (bool, error) {
	err := InitMap_Name_Pwd()
	if err != nil {
		return false, err
	}
	mp := Map_Name_Pwd

	//判断用户是否存在
	pwdBase := mp[name]
	if pwdBase == "" {
		str := fmt.Sprintf("用户:%v不存在!",name)
		return false, errors.New(str)
	}

	//对用户的密码pwd进行hash和base编码
	pwdBytes := Hash.Sha256Hash([]byte(pwd))
	pwdBase64Bytes := Base.Base64Encode(pwdBytes)
	pwdBaseStr := string(pwdBase64Bytes)

	//比较密码是否相等并返回结果
	return pwdBase == pwdBaseStr, nil
}

//根据邮箱登入
func QueryByMail(mailBox string, pwd string) (bool, error) {
	err := InitMap_MailBox_Pwd()
	if err != nil {
		return false, err
	}
	mp := Map_MailBox_Pwd

	//判断邮箱是否存在
	pwdBase := mp[mailBox]
	if pwdBase == "" {
		str := fmt.Sprintf("邮箱:%v不存在!",mailBox)
		return false, errors.New(str)
	}
	//对用户的密码pwd进行hash和base编码
	pwdBytes := Hash.Sha256Hash([]byte(pwd))
	pwdBase64Bytes := Base.Base64Encode(pwdBytes)
	pwdBaseStr := string(pwdBase64Bytes)

	//比较密码是否相等并返回结果
	return pwdBase == pwdBaseStr, nil
}

//保存用户信息
func SaveUserInfo(name, mailBox, pwd string) error {

	if name == "" || mailBox == "" {
		return errors.New("用户名和邮箱不能为空!")
	}

	//0.判断用户名和邮箱是否已存在
	err := InitMap_Name_Pwd()
	if err != nil {
		return err
	}
	if Map_Name_Pwd[name] != "" {
		str := fmt.Sprintf("注册失败,用户名:%v 已存在!", name)
		return errors.New(str)
	}

	err = InitMap_MailBox_Name()
	if err != nil {
		return err
	}

	if Map_MailBox_Name[mailBox] != "" {
		str := fmt.Sprintf("注册失败,邮箱:%v 已注册!", mailBox)
		return errors.New(str)
	}

	//1.保存用户名和密码
	err = SaveNamePwd(name, pwd, mailBox)
	if err != nil {
		return errors.New("保存用户名和密码出现错误，请重试！")
	}

	//2.保存用户名和密码
	err = SaveMailBoxPwd(mailBox, pwd, name)
	if err != nil {
		return err
	}

	//3.保存用户邮箱和用户名
	err = SaveMailBoxName(mailBox, name)
	if err != nil {
		return err
	}

	//4.保存用户名和邮箱
	err = SaveNameMailBox(name, mailBox)
	if err != nil {
		return err
	}

	return nil
}

//保存用户名和密码
func SaveNamePwd(name, pwd, mailBox string) error {
	//1.打开文件name_Pwd.txt，保存用户名和密码
	f_name_Pwd, err := os.OpenFile(utils.Name_Pwd_PATH, os.O_WRONLY, 0644) //传递文件路径
	if err != nil {
		fmt.Println("err = ", err, )
		return err
	}
	name_Pwd_n, err := f_name_Pwd.Seek(0, 2)
	if err != nil {
		return err
	}
	//使用完毕，需要关闭文件
	defer f_name_Pwd.Close()

	err = InitMap_Name_Pwd()
	if err != nil {
		return err
	}

	if Map_Name_MailBox[name] != "" {
		str := fmt.Sprintf("注册失败,邮箱:%v 已存在!", Map_Name_MailBox[name])
		return errors.New(str)
	}
	pwdBytes := Hash.Sha256Hash([]byte(pwd))
	base64PwdBytes := Base.Base64Encode(pwdBytes)
	pwdStr := string(base64PwdBytes)
	a := name + "," + pwdStr + "\n"
	_, err = f_name_Pwd.WriteAt([]byte(a), name_Pwd_n)
	if err != nil {
		return err
	}

	err = InitMap_Name_Pwd()
	if err != nil {
		return err
	}
	return nil
}

//保存用户邮箱和密码
func SaveMailBoxPwd(mailBox, pwd, mame string) error {
	f_mailBox_Pwd, err := os.OpenFile(utils.MailBox_Pwd_PATH, os.O_WRONLY, 0644) //传递文件路径
	if err != nil {
		fmt.Println("err = ", err, )
		return err
	}
	mailBox_Pwd_n, err := f_mailBox_Pwd.Seek(0, 2)
	if err != nil {
		return err
	}
	//使用完毕，需要关闭文件
	defer f_mailBox_Pwd.Close()

	err = InitMap_MailBox_Pwd()
	if err != nil {
		return err
	}

	//if Map_Name_Pwd[mailBox] != "" {
	//	str := fmt.Sprintf("注册失败,邮箱:%v 已存在!", mailBox)
	//	return errors.New(str)
	//}
	pwdBytes := Hash.Sha256Hash([]byte(pwd))
	base64PwdBytes := Base.Base64Encode(pwdBytes)
	pwdStr := string(base64PwdBytes)
	a := mailBox + "," + pwdStr + "\n"
	_, err = f_mailBox_Pwd.WriteAt([]byte(a), mailBox_Pwd_n)
	if err != nil {
		return err
	}

	err = InitMap_MailBox_Pwd()
	if err != nil {
		return err
	}

	return nil
}

//保存用户名和邮箱
func SaveNameMailBox(name, mailBox string) error {
	f_Name_MailBox, err := os.OpenFile(utils.Name_MailBox_PATH, os.O_WRONLY, 0644) //传递文件路径
	if err != nil {
		return err
	}
	Name_MailBox_n, err := f_Name_MailBox.Seek(0, 2)
	if err != nil {
		return err
	}
	//使用完毕，需要关闭文件
	defer f_Name_MailBox.Close()

	err = InitMap_Name_MailBox()
	if err != nil {
		return err
	}

	a := name + "," + mailBox + "\n"
	_, err = f_Name_MailBox.WriteAt([]byte(a), Name_MailBox_n)
	if err != nil {
		return err
	}

	err = InitMap_Name_MailBox()
	if err != nil {
		return err
	}

	return nil
}

//保存用户邮箱和用户名
func SaveMailBoxName(mailBox, name string) error {
	f_MailBox_Name, err := os.OpenFile(utils.MailBox_Name_PATH, os.O_WRONLY, 0644) //传递文件路径
	if err != nil {
		return err
	}
	MailBox_Name_n, err := f_MailBox_Name.Seek(0, 2)
	if err != nil {
		return err
	}
	//使用完毕，需要关闭文件
	defer f_MailBox_Name.Close()

	err = InitMap_MailBox_Name()
	if err != nil {
		return err
	}

	a := mailBox + "," + name + "\n"
	_, err = f_MailBox_Name.WriteAt([]byte(a), MailBox_Name_n)
	if err != nil {
		return err
	}

	err = InitMap_MailBox_Name()
	if err != nil {
		return err
	}
	return nil
}

//初始化Map_Name_Pwd
func InitMap_Name_Pwd() error {
	Map_Name_Pwd = make(map[string]string)

	f, err := os.Open("db/name_Pwd.txt")
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
			if i == 1{
				if Map_Name_Pwd[br[0]] == "" {
					Map_Name_Pwd[br[0]] = br[i]
				}
			}
		}

	}
	return nil
}

//初始化Map_MailBox_Pwd
func InitMap_MailBox_Pwd() error {
	Map_MailBox_Pwd = make(map[string]string)

	f, err := os.Open("db/mailBox_Pwd.txt")
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
			if i == 1{
				if Map_MailBox_Pwd[br[0]] == "" {
					Map_MailBox_Pwd[br[0]] = br[i]
				}
			}
		}

	}
	return nil
}

//初始化Map_Name_MailBox
func InitMap_Name_MailBox() error {
	Map_Name_MailBox = make(map[string]string)

	f, err := os.Open("db/name_MailBox.txt")
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
			if i == 1{
				if Map_Name_MailBox[br[0]] == "" {
					Map_Name_MailBox[br[0]] = br[i]
				}
			}
		}

	}
	return nil
}

//初始化Map_MailBox_Name
func InitMap_MailBox_Name() error {
	Map_MailBox_Name = make(map[string]string)

	f, err := os.Open("db/mailBox_Name.txt")
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
			if i == 1{
				if Map_MailBox_Name[br[0]] == "" {
					Map_MailBox_Name[br[0]] = br[i]
				}
			}
		}

	}
	return nil
}