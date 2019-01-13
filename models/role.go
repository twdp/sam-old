package models

import "tianwei.pro/business/model"

// 部门角色
type Role struct {

	model.Base

	Name string

	Status int8

	// 如果id=0并且系统是不拉平数据角色权限，那么这个角色就是模板角色
	BranchId int64

	// 这个角色是从哪个模板角色上映射出来的
	FromId int64



}