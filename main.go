package main

import (
	"page/model"
	"page/router"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	model.Register()
}

func main() {
	router.Register()
	beego.Run()
}
