package models

import (
	"github.com/astaxie/beego/orm"
	"tianwei.pro/business/model"
)

const (
	Menu = iota
	Page
	Button
)

// todo:: 前端选择资源时，勾选的是哪一个框，就给后端哪一个框
// 后端需要根据选择的资源，向上递归，找出页面和按钮
type Api struct {

	model.Base

	// 展示的名称
	Name string

	// url路径
	Path string `orm:"size(100)"`

	// PUT/GET/DELETE/POST
	Method string `orm:"size(20)"`

	// 菜单和页面使用的那妞
	Icon string

	// 菜单排序
	// 从小到大展示
	Order int64

	// api 类型
	// Menu   菜单
	// Page   页面
	// Button  按钮
	Type int8

	// 菜单->页面->按钮
	ParentId int64

	// 本api来自哪个系统
	SystemId int64

	// api当前状态
	// 前端获取api状态，即可实现动态api更换
	// 可以实现发布后，快速回滚
	Status int8
}

// 多字段唯一键
func (a *Api) TableUnique() [][]string {
	return [][]string{
		{ "Path", "Method", "BranchId", },
	}
}

func init() {
	orm.RegisterModelWithPrefix("sam_", new(Api))
}