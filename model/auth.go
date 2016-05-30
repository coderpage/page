package model

import "time"

type Auth struct {
	Id         int64     // 自增 id
	Uid        int64     // user id
	Token      string    `orm:"size(64)"` // 令牌
	Server     string    // 授权网站
	Status     string    // 状态
	Type       string    // 授权类型
	Redirect   string    // 回调
	ExpiryDate time.Time // 有效期
}
