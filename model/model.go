package model

import (
	"page/conf"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(User), new(Auth))
}

// 注册数据库表
func Register() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	mysqlUser := conf.MysqlUser
	mysqlDb := conf.MysqlDBName
	mysqlPwd := conf.MysqlPass

	//tcp(112.80.45.162:9099)
	orm.RegisterDataBase("default", "mysql", mysqlUser+":"+mysqlPwd+"@/"+mysqlDb+"?charset=utf8&loc=Local")

	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
}
