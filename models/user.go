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
	err := db.SaveUserInfo(u.UserName, u.MailBox, u.Password)
	if err != nil {
		return err
	}
	return nil
}
//用于查询用户信息
func (u User) QueryUser() (bool,error) {
	fmt.Println("查询",u.Password)
	fmt.Println("查询",u.UserName)
	return db.QueryByName(u.UserName,u.Password)
}

