package storage

import (
	"page/model"

	"github.com/astaxie/beego/orm"
)

const (
	authKeyToken string = "Token"
)

// AddNewAuth 保存一个新的 Auth
func AddNewAuth(auth *model.Auth) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(auth)
	return id, err
}

// FindAuthByToken 通过 Token 查询 Auth
func FindAuthByToken(token string) (auth *model.Auth, err error) {
	auth = &model.Auth{Token: token}
	o := orm.NewOrm()
	err = o.QueryTable(auth).Filter(authKeyToken, token).One(auth)

	return auth, err
}
