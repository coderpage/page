package model

import (
	"encoding/json"
	"time"
)

// User 用户表
type User struct {
	Id          int64     `json:"id"`                              // 主键id
	UserName    string    `json:"user_name" orm:"size(100)"`       // 用户名
	Password    string    `json:"-"`                               // 用户密码
	Email       string    `json:"email"`                           // 用户邮箱
	DisplayName string    `json:"display_name"`                    // 用户显示的名称
	Url         string    `json:"url"`                             // 用户主页
	Created     time.Time `json:"created" orm:"auto_now_add"`      // 用户注册的时间
	Activated   time.Time `json:"activated"`                       // 用户最后活动的时间
	Logged      time.Time `json:"logged"`                          // 用户上次登录的时间
	Group       string    `json:"group" orm:"default(nonactived)"` // 用户组 nonactived | common | admin
}

// 返回 user 转为字符串格式数据，默认为 json 格式
func (this *User) String() string {
	// str := fmt.Sprintf("User: {Id:%d  UserName:%s  Password:%s  Email:%s  DisplayName:%s  Group:%s}", this.Id, this.UserName, this.Password, this.Email, this.DisplayName, this.Group)
	return this.JsonString()
}

// 返回 user 的 json 格式字符串
func (this *User) JsonString() string {
	jbytes, err := json.Marshal(this)
	if err != nil {
		return "{}"
	}
	return string(jbytes)
}
