package models

import "time"

type User struct {
	UserId int `json:"user_id" db:"user_id"`
	UserName string `json:"user_name" db:"user_name"`
	UserPhone string `json:"user_phone" db:"user_phone"`
	UserBalance float32 `json:"user_balance" db:"user_balance"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
}
