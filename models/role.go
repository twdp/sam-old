package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"tianwei.pro/business/model"
)

// 部门角色
type Role struct {

	model.Base

	Name string

	Status int8

	// 角色属于哪个系统
	SystemId int64

	// 如果id=0并且系统是不拉平数据角色权限，那么这个角色就是模板角色
	BranchId int64

	// 这个角色是从哪个模板角色上映射出来的
	FromId int64

	// 本角色拥有的权限集
	PermissionSet string `orm:"type(text)" json:"-"`
}

// 多字段唯一键
func (r *Role) TableUnique() [][]string {
	return [][]string{
		{ "Name", "BranchId", "SystemId", },
	}
}

func init() {
	orm.RegisterModelWithPrefix("sam_", new(Role))
}

func LoadByRolesAndStatus(systemId int64, status int8) ([]*Role, error) {
	var roles []*Role
	if _, err := orm.NewOrm().QueryTable(&Role{}).Filter("SystemId", systemId).Filter("Status", status).All(&roles); err != nil {
		logs.Error("load by roleIds and systemId and status failed. roleIds: %v, systemId: %d, status: %d, err: %v", roles, systemId, status, err)
		return nil, errors.New("查询用户角色失败")
	} else {
		return roles, nil
	}
}