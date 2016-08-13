package model

import (
	"time"
)

type Article struct {
	Id           int64     `json:"id"`                               // 文章表主键id
	UserId       int64     `json:"user_id"`                          // 文章所属用户id
	Title        string    `json:"title"`                            // 文章标题
	Created      time.Time `json:"created" orm:"index;auto_now_add"` // 文章创建时间
	Updated      time.Time `json:"updated" orm:"index;auto_now"`     // 文章修改时间
	Content      string    `json:"content" orm:"type(text);null"`    // 文章内容
	Type         string    `json:"type"`                             // 文章类别
	Status       string    `json:"status"`                           // 文章状态 （发布、草稿）publish || draft
	CommentsNum  int       `json:"comment_num"`                      // 评论数
	ViewsNum     int64     `json:"view_num"`                         // 浏览数
	AllowComment bool      `json:"allow_comment" orm:"default(1)"`   // 是否允许评论
	AllowPing    bool      `json:"allow_ping" orm:"default(1)"`      // 是否允许 ping
	AllowFeed    bool      `json:"allow_feed" orm:"default(1)"`      // 是否允许出现在聚合中
	Parent       int64     `json:"parent"`                           // parent
}

func NewArticle(title string, content string) (article *Article) {
	return &Article{Title: title, Content: content}
}
