package router

import (
	"page/router/api"
)

// Register 注册路由
func Register() {
	registerAuths()
	registerCommons()
	api.RegisterArticles()

}
