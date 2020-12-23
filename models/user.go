package models

import (
	"bitcoin-Core/db"
	"bitcoin-Core/utils"
	"fmt"
)

type User struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

//用于保存用户信息
func (u User) SaveUser() error {
	fmt.Println("注册",u.UserName)
	fmt.Println("注册",u.Password)
	return db.SaveUser(utils.PATH, u.UserName, u.Password)
}
//用于查询用户信息
func (u User) QueryUser() (bool,error) {
	fmt.Println("查询",u.Password)
	fmt.Println("查询",u.UserName)
	return db.Query(u.UserName,u.Password)
}

