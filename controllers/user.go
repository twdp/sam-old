package controllers

import (
	"fmt"
	"strings"
	"tianwei.pro/business"
	"tianwei.pro/business/controller"
	"tianwei.pro/sam/models"
)

const (
	base int64 = 1
	position = base << 6
)

type UserController struct {
	controller.RestfulController
}

// 验证app key
// 获取用户信息
// 根据系统信息和用户信息查找对应的角色列表
// 是否需要查询数据权限信息，查出数据权限信息
// 根据角色信息找出对应的菜单、按钮、页面信息

// role -> {
// 	branchTree: {
//
//  }
//  roleId:
//  roleName: ,
//  permissionUrls: [
//   '/api/v1...'
//  ],
//  menuTrees: [
//  	MenuTree
//  ]
// }
// MenuTree: {
// 		menu: ,
//      type: ,
//      menus: [ MenuTree ]
// }
// @router /permission [get]
func (u *UserController) LoadPermission() {
	appKey := u.GetString("app")
	uid := u.GetSession(UserSessionName).(*models.User).Id
	s := &models.System{ AppKey: appKey, }
	if err := s.FindByAppKey(); err != nil {
		u.E500(err.Error())
	}

	var roles []*models.Role

	userRole := &models.UserRole{ UserId: uid, SystemId: s.Id }
	if uroles, err := userRole.LoadByUserAndSystemId(); err != nil {
		u.E500(err.Error())
	} else if len(uroles) == 0 {
		u.ReturnJson([]string{})
	} else {
		var roleIds []int64
		for _, role := range uroles {
			roleIds = append(roleIds, role.RoleId)
		}
		if rr, err := models.LoadByRoleIdsAndSystemIdAndStatus(roleIds, s.Id, models.Active); err != nil {
			u.E500(err.Error())
		} else {
			roles = rr
		}
	}

	type UrlMap struct {
		Path string
		Method string
		PermissionSet string `json:"-"`
	}

	type Response struct {

		RoleName string `json:"role_name"`
		RoleId int64 `json:"role_id"`
		PermisisionUrls []*UrlMap `json:"permisision_urls"`
	}
	urlIds := make(map[int64]*UrlMap)

	if apis, err := models.LoadApiBySystemAndStatus(s.Id, models.Active); err != nil {
		u.E500(err.Error())
	} else {
		for _, api  := range apis {
			ids := strings.Split(api.ReplaceIds, ",")
			urlIds[api.Id] = &UrlMap{
				Path: api.Path,
				Method: api.Method,
				PermissionSet: api.PermissionSet,
			}
			for _, id := range ids  {
				if id == "" {
					continue
				}
				urlIds[business.CastStringToInt64(id)] = &UrlMap{
					Path: api.Path,
					Method: api.Method,
					PermissionSet: api.PermissionSet,
				}
			}
		}
	}



	for _, role := range roles {
		ps := strings.Split(role.PermissionSet, ",")
		var per []*UrlMap

		for index, p := range ps {
			pp := business.CastStringToInt64(p)

			var ii uint = 0
			for pp > 0 {
				if pp % 2 == 1 {
					id := business.CastIntToInt64(index) * position + base << ii
					if urlMap, exist := urlIds[id]; exist {
						per = append(per, urlMap)
					}
				}
				ii++
				pp = pp / 2
			}
		}

		r := &Response{
			RoleId: role.Id,
			RoleName: role.Name,
			PermisisionUrls:per,
		}
		fmt.Println(r)
	}

	u.ReturnJson([]string{})

}