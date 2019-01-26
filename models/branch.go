package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"tianwei.pro/business/model"
)

// 组织架构表
type Branch struct {

	model.Base

	// 名称
	Name string

	// 父id
	ParentId int64
}

// 多字段唯一键
func (b *Branch) TableUnique() [][]string {
	return [][]string{
		{ "Name", "ParentId", },
	}
}

func init() {
	orm.RegisterModelWithPrefix("sam_", new(Branch))
}

func FindByPid(pid int64) ([]*Branch, error) {
	var childrens []*Branch
	if _, err := orm.NewOrm().QueryTable(&Branch{}).Filter("ParentId", pid).All(&childrens); err != nil {
		logs.Error("branch find by pid failed. err: %v", err)
		return nil, err
	}
	return childrens, nil
}