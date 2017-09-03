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

// FindAuthByKey 通过 key 查询 Auth
func FindAuthByKey(key string) (auth *model.Auth, err error) {
	auth = &model.Auth{Key: key}
	o := orm.NewOrm()
	err = o.QueryTable(auth).Filter("key", key).One(auth)

	return auth, err
}

// FindAuthByKeyLatest 通过 key 查询最新的 Auth
func FindAuthByKeyLatest(key string) (auth *model.Auth, err error) {
	auth = &model.Auth{Key: key}
	o := orm.NewOrm()
	err = o.QueryTable(auth).Filter("key", key).OrderBy("-id").One(auth)

	return auth, err
}
