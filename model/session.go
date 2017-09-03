package model

import (
	"time"
)

const (
	TableName_Session = "session"
)

type Session struct {
	Token      string `orm:"pk"`
	ExpiryDate time.Time
	UserId     int64
	UserName   string
	UserInfo   string
}
