package storage

import (
	"github.com/astaxie/beego/orm"
	"page/model"
	"page/tool/secure"
	"time"
)

const (
	userKeyID    string = "Id"
	userKeyEmail string = "Email"
	userKeyPwd   string = "Password"
)

// CreateUser create a new user in mysql, if user is allready exist in table `user`,
// or some errors occurred during creating user, return error
func CreateUser(user *model.User) error {
	// if user is nil, return error
	if user == nil {
		return ErrIllegalArg
	}
	// if email or pwd is empty, return error
	if user.Email == "" || user.Password == "" {
		return ErrIllegalArg
	}

	encryptoPass, err := secure.DoMd5(user.Password)
	if err != nil {
		return ErrInternal
	}

	// init default UserName DisplayName by Email
	user.Password = encryptoPass
	user.UserName, user.DisplayName, user.Created, user.Activated, user.Logged, user.Group =
		user.Email, user.Email, time.Now(), time.Now(), time.Now(), model.UserGroupNoActived

	// save user to mysql table `user`
	o := orm.NewOrm()
	created, id, err := o.ReadOrCreate(user, "Email")

	if err != nil {
		return err
	}

	if created {
		user.Id = id
		return nil
	}

	return ErrRowExist
}

// DeleteUser delete user in mysql table `user`, user.Id must be setted
func DeleteUser(user *model.User) error {
	if user == nil {
		return ErrIllegalArg
	}

	if user.Id == 0 {
		return ErrIllegalArg
	}

	o := orm.NewOrm()
	_, err := o.Delete(user)
	if err != nil {
		return err
	}

	return err
}

// FindUserById 通过 User:Id 查询 User
func FindUserById(uid int64) (user *model.User, err error) {
	o := orm.NewOrm()
	user = &model.User{Id: uid}
	err = o.Read(user, userKeyID)
	if err != nil {
		return nil, ErrNoRows
	}
	return user, err
}

// FindUserByEmail 通过 User:Email 查询 User
func FindUserByEmail(email string) (user *model.User, err error) {
	o := orm.NewOrm()
	user = &model.User{Email: email}
	err = o.Read(user, userKeyEmail)
	if err != nil {
		return nil, ErrNoRows
	}

	return user, nil
}

func UpdateUser(user *model.User, columns ...string) (err error) {
	o := orm.NewOrm()

	if user != nil && user.Id == 0 {
		return ErrIllegalArg
	}

	_, err = o.Update(user, columns...)

	return err
}

// CheckEmailPwd  check user email and password in mysql table user
func CheckEmailPwd(user *model.User) (err error) {
	if user == nil {
		return ErrIllegalArg
	}

	if user.Email == "" || user.Password == "" {
		return ErrIllegalArg
	}

	encryptoPass, err := secure.DoMd5(user.Password)
	if err != nil {
		return ErrInternal
	}
	user.Password = encryptoPass
	o := orm.NewOrm()
	err = o.Read(user, userKeyEmail, userKeyPwd)

	if err == orm.ErrNoRows {
		return ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}
