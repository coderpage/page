package storage

import (
	"blog/model"

	"github.com/astaxie/beego/orm"
)

func SaveSession(sess *model.Session) error {
	o := orm.NewOrm()

	_, _, err := o.ReadOrCreate(sess, "Token")
	return err
}

func DeleteSessionByToken(tk string) error {
	o := orm.NewOrm()

	_, err := o.QueryTable("session").Filter("token", tk).Delete()
	return err
}

func UpdateSession(see *model.Session, cols ...string) error {
	o := orm.NewOrm()

	_, err := o.Update(see, cols...)
	return err
}

func GetSessionByToken(tk string) (sess *model.Session, err error) {
	o := orm.NewOrm()
	sess = new(model.Session)
	err = o.QueryTable("session").Filter("Token", tk).One(sess)
	return sess, err
}
