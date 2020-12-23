package models

import (
	"BWP/dbMysql"
	"fmt"
)

type User struct {
	Id int `form:"id"`
	Name string `form:"name"`
	Password string `form:"password"`
}

func (u User) Queryuser () (*User , error) {
	row := dbMysql.DB.QueryRow("select * from user where user_name = ? and user_pwd = ?",u.Name,u.Password)
	err := row.Scan(&u)
	fmt.Println(u)
	if err != nil {
		return nil , err
	}
	return &u , nil
}