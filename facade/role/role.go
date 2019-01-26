package role

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"tianwei.pro/business"
	"tianwei.pro/sam/facade/cache"
	"tianwei.pro/sam/models"
	"time"
)

var (
	roleCache = cache.NewCache()

	userRoleCache = cache.NewCache()

	InvalidRoleCache     = errors.New("删除角色缓存失败")
	InvalidUserRoleCache = errors.New("删除用户角色关联关系失败")
)

func LoadRolesBySystemId(systemId int64) ([]*models.Role, error) {
	sid := business.CastInt64ToString(systemId)
	if roles := roleCache.Get(sid); roles != nil {
		return roles.([]*models.Role), nil
	} else {
		if roles, err := models.LoadByRolesAndStatus(systemId, models.Active); err != nil {
			return nil, err
		} else {
			roleCache.Put(sid, roles, time.Duration(30)*time.Minute)
			return roles, nil
		}
	}
}

func LoadRolesByUIdAndSID(userId, sid int64) ([]*models.UserRole, error) {
	uid := fmt.Sprintf("%d_%d", userId, sid)
	if roles := userRoleCache.Get(uid); roles != nil {
		return roles.([]*models.UserRole), nil
	} else {
		userRole := &models.UserRole{UserId: userId, SystemId: sid}
		if roles, err := userRole.LoadByUserAndSystemId(); err != nil {
			return nil, err
		} else {
			roleCache.Put(uid, roles, time.Duration(30)*time.Minute)

			return roles, nil
		}
	}
}

func LoadRoleIdsByUId(userId, sid int64) ([]int64, error) {
	if roles, err := LoadRolesByUIdAndSID(userId, sid); err != nil {
		return nil, err
	} else {
		var roleIdss []int64

		for _, role := range roles {
			roleIdss = append(roleIdss, role.Id)
		}
		return roleIdss, nil
	}
}

// 清空用户角色
func InvalidByUIdAndSystemId(uid, sid int64) error {
	key := fmt.Sprintf("%d_%d", uid, sid)
	if err := userRoleCache.Delete(key); err != nil {
		logs.Error("invalid user role cache failed. uid: %d, sid: %d, err： %v", uid, sid, err)
		return InvalidUserRoleCache
	}
	return nil
}

// 清除角色缓存
func InvalidBySystemId(systemId int64) error {
	if err := roleCache.Delete(business.CastInt64ToString(systemId)); err != nil {
		logs.Error("invalid cache failed. systemId: %d", systemId)
		return InvalidRoleCache
	}
	return nil
}
