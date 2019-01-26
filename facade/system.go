package facade

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"tianwei.pro/business"
	cache2 "tianwei.pro/sam/facade/cache"
	"tianwei.pro/sam/models"
	"time"
)
var (
	keyCache = cache2.NewCache()

	pidCache = cache2.NewCache()

	LoadSystemError = errors.New("查询系统失败")
)

// 根据系统的 app key查询系统信息
func FindByKey(key string) (*models.System, error) {
	system := keyCache.Get(key)
	if system == nil {
		system := &models.System{
			AppKey: key,
		}
		if err := system.FindByAppKey(); err != nil {
			logs.Error("system find by app key failed. appkey: %s, err: %v", key, err)
			return nil, LoadSystemError
		}
		return system, nil
	} else {
		return system.(*models.System), nil
	}

}


func FindByPid(pid int64) ([]int64, error) {
	var childrens []*models.Branch
	if cacheChildrens := pidCache.Get(business.CastInt64ToString(pid)); cacheChildrens == nil {
		if childrenss, err := models.FindByPid(pid); err != nil {
			return nil, err
		} else {
			pidCache.Put(business.CastInt64ToString(pid), childrens, time.Duration(30) * time.Minute)
			childrens = childrenss
		}
	} else {
		childrens = cacheChildrens.([]*models.Branch)
	}

	var branchIds []int64
	for _, children := range childrens {
		branchIds = append(branchIds, children.Id)
		if posterities, err := FindByPid(children.Id); err != nil {
			return nil, err
		} else {
			branchIds = append(branchIds, posterities...)
		}
	}
	return branchIds, nil
}
