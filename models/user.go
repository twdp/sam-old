package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"tianwei.pro/business"
	"tianwei.pro/business/model"
)

var (
	UserNotFound = errors.New("账号或密码错误")
	SystemError = errors.New("获取用户信息失败")
)
type User struct {

	model.Base

	UserName string `orm:"size(64);unique"`

	DisplayName string `orm:"size(64)"`

	Avatar string

	Email string `orm:"size(64);unique"`

	Phone string `orm:"size(64);unique"`

	Sex int8

	Password string

	Type int8

	Status int8
}


func init() {
	orm.RegisterModelWithPrefix("sam_", new(User))
}

func (u *User) FindByEmail() error {
	if err := orm.NewOrm().Read(u, "Email"); err != nil {
		if business.IsNoRowsError(err) {
			return UserNotFound
		}
		logs.Error("find user by email failed. user: %v, err: %v", u, err)
		return SystemError
	}
	return nil
}

