package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"tianwei.pro/business"
	"tianwei.pro/business/controller"
	"tianwei.pro/sam/facade"
	"tianwei.pro/sam/models"
)

type PortalController struct {
	controller.RestfulController
}

const (
	UserSessionName = "_sam_user_info"
)

type SamClaims struct {
	jwt.StandardClaims
	UserName string `json:"user_name;omitempty"`
	Email string `json:"email;omitempty"`
	Phone string `json:"phone;omitempty"`
	Id int64 `json:"id;omitempty"`
	Avatar string `json:"avatar;omitempty"`
}

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
		user.Password = ""
		u.SetSession(UserSessionName, user)

		u.SetSecureCookie(beego.AppConfig.DefaultString("tokenSecret", "__sam__"), "_sam_token_", token, beego.AppConfig.DefaultInt64("tokenExpire", 24 * 60 * 30) * 3600, "", "", "", true)
		u.ReturnJson(map[string]interface{} {
			"id": user.Id,
			"user_name": user.UserName,
			"email": user.Email,
			"phone": user.Phone,
			"token": token,
		})
	}
}

// @router /logout [post]
func (u *PortalController) Logout() {
	u.SetSecureCookie(beego.AppConfig.DefaultString("tokenSecret", "__sam__"), "_sam_token_", "", -1)

}