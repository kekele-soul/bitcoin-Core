package models

import (
	"bitcoin-Core/db"
	"fmt"
)

type User struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func (u User) SaveUser() error {
	fmt.Println("注册",u.UserName)
	fmt.Println("注册",u.Password)
	return db.SaveUser("db/user.txt", u.UserName, u.Password)
}
func (u User) QueryUser() (bool,error) {
	fmt.Println("查询",u.Password)
	fmt.Println("查询",u.UserName)
	return db.Query(u.UserName,u.Password,"db/user.txt")
}

