package router

import (
	"page/router/api"
	"page/router/web"
)

// Register 注册路由
func Register() {
	registerAuths()
	registerCommons()
	api.RegisterArticles()
	web.Register()
}
