package models

import (
	"bitcoin-Core/db"
)

type User struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func (u User) SaveUser() error {
	return db.SaveUser("db/user.txt", u.UserName, u.Password)
}


