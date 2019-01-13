package models

import "tianwei.pro/business/model"

// 对接的系统
type System struct {

	model.Base

	Name string

	AppKey string

	Secret string

	// 本系统是否使用数据权限
	UseDataPermission bool

	// 数据权限和角色是否拉平
	Lateral bool


	// 使用了模板角色，是否对模板角色可见
	TemplateRoleVisible bool

	// todo:: 使用模板角色，拷贝一份出来，然后让其可修改
	// 当前：要么完全自己设置角色权限，要么使用模板角色，但不允许修改
	// 系统对接设置以后，尽量不能修后面这3个字段，如果需要修改，需要清除所有设置
}

