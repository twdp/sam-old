package models

import (
	"github.com/astaxie/beego/orm"
	"tianwei.pro/business/model"
)

// 功能点表
// 多个api中的url关联城功能点
// 权限树选择的时候可以反关联各个url
// todo:: 页面上添加api的时候，选择是否是功能点
// 然后可以对功能点添加api，进行功能点聚合
type Function struct {
	model.Base

	// 功能点名称
	Name string `orm:"size(64)"`

	// 主id
	MasterId int64 `orm:"unique"`

	// 权限集
	PermissionSet string `orm:"size(100);index"`
}

// e.g.  创建-> 查看、修改
//       修改-> 查看
//       查看  只有查看权限

func init() {
	orm.RegisterModelWithPrefix("sam_", new(Function))
}