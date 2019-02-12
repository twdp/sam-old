package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"tianwei.pro/business"
	"tianwei.pro/business/controller"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam/facade"
	"tianwei.pro/sam/models"
	"tianwei.pro/sam/upper"
)

// portal管理接口
type PortalController struct {
	controller.RestfulController
}

// @Title 邮箱登录接口
// @Description 根据邮箱和密码进行登录
// @Param email query string true "需要登录的邮箱"
// @Param password query string true "密码"
// @Success 200 {object} map
// @Failure 500 {string} 错误信息
// @router /login-by-email [post]
func (u *PortalController) LoginByEmail() {
	email := u.GetString("email")
	pass := u.GetString("password")

	user := &models.User{ Email: email }
	if err := user.FindByEmail(); err != nil {
		u.E500(err.Error())
	}

	if _, err := business.ValidateCrypto(pass, user.Password); err != nil {
		logs.Error("bcrypt compare hash failed. pass: %s, user: %v, err: %v", pass, user, err)
		u.E500("账号或密码错误")
	}

	if token, err := facade.T.EncodeToken(user); err != nil {
		u.E500(err.Error())
	} else {
		u.SetSecureCookie(beego.AppConfig.DefaultString("tokenSecret", "__sam__"), "_sam_token_", token, beego.AppConfig.DefaultInt64("tokenExpire", 24 * 60 * 30) * 3600, "", "", "", true)

		agentUserInfo := &sam_agent.UserInfo{}
		param := &sam_agent.VerifyTokenParam{
			SystemInfoParam: sam_agent.SystemInfoParam{
				AppKey: beego.AppConfig.String("appKey"),
				Secret: beego.AppConfig.String("secret"),
			},
		}
		if err := upper.SamAgentFacade.VerifyToken(context.Background(), param, agentUserInfo); err != nil {
			u.E500("系统错误")
			return
		} else {
			u.SetSession(sam_agent.SamUserInfoSessionKey, agentUserInfo)
		}
		u.ReturnJson(map[string]interface{} {
			"id": user.Id,
			"user_name": user.UserName,
			"email": user.Email,
			"phone": user.Phone,
			"token": token,
		})
	}
}