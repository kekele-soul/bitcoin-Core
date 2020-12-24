package models

import (
	"bitcoin-Core/db"
	"fmt"
)

type User struct {
	UserName string `form:"username"`
	Password string `form:"password"`
	MailBox string `form:"email"`//邮箱
}

//用于保存用户信息
func (u User) SaveUser() error {
	fmt.Println("注册",u.UserName)
	fmt.Println("注册",u.Password)
	fmt.Println("注册",u.MailBox)
	err := db.SaveUserInfo(u.UserName, u.MailBox, u.Password)
	if err != nil {
		return err
	}
	return nil
}
//用于查询用户账号信息
func (u User) QueryUserByUserName() (bool,error) {
	fmt.Println("查询",u.Password)
	fmt.Println("查询",u.UserName)
	return db.QueryByName(u.UserName,u.Password)
}
//用于查询用户邮箱信息
func (u User) QueryUserByEmail() (bool,error) {
	fmt.Println("查询",u.Password)
	fmt.Println("查询",u.MailBox)
	return db.QueryByMail(u.MailBox,u.Password)
}

