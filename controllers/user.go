package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"tianwei.pro/business"
	"tianwei.pro/business/controller"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam/facade"
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
	if u.GetSession(sam_agent.SamUserInfoSessionKey) == nil {
		u.E500("请重新登录")
		return
	}
	uid := u.GetSession(sam_agent.SamUserInfoSessionKey).(*sam_agent.UserInfo).Id
	s := &models.System{ AppKey: appKey, }
	if err := s.FindByAppKey(); err != nil {
		u.E500(err.Error())
	}

	var roles, err = facade.FindOwnRoles(uid, s.Id)
	if err != nil {
		u.E500(err.Error())
		return
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


	var responses []*Response

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
		responses = append(responses, r)
	}

	u.ReturnJson(responses)

}


// @router /logout [post]
func (u *PortalController) Logout() {
	u.SetSecureCookie(beego.AppConfig.DefaultString("tokenSecret", "__sam__"), "_sam_token_", "", -1)
}