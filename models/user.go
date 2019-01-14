package models

import (
	"github.com/astaxie/beego/orm"
	"tianwei.pro/business/model"
)

type User struct {

	model.Base

	UserName string `orm:"size(64);unique"`

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