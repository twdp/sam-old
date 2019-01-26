package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"tianwei.pro/business/model"
)

type UserRole struct {
	model.Base

	UserId int64

	RoleId int64

	SystemId int64

	// 如果系统角色和数据权限拉平
	// 数据权限放到此字段中
	// 前端存过来的都是选中的那一级
	BranchIds string
}

func init() {
	orm.RegisterModelWithPrefix("sam_", &UserRole{})
}

func (u *UserRole) LoadByUserAndSystemId() ([]*UserRole, error) {
	var roles []*UserRole
	if _, err := orm.NewOrm().QueryTable(u).Filter("UserId", u.UserId).Filter("SystemId", u.SystemId).All(&roles); err != nil {
		logs.Error("load user role by userAndSystemId failed. userRole: %v, err: %v", u, err)
		return nil, errors.New("查询用户权限失败")
	}
	return roles, nil
}