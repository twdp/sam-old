package controllers

import (
	"fmt"
	"golang.org/x/crypto/md4"
	"sync"
	"tianwei.pro/business"
	"tianwei.pro/business/controller"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam/models"
	"time"
)

type SystemController struct {
	sync.Mutex
	controller.RestfulController
}

type SystemVO struct {
	Name                string `valid:"Required`
	KeepSign            bool
	Strategy            int8
	TemplateRoleVisible bool
}

// @router / [post]
func (s *SystemController) AddSystem() {
	userInfo := s.GetSession(sam_agent.SamUserInfoSessionKey).(*sam_agent.UserInfo)

	svo := &SystemVO{}
	if err := s.ReadBody(svo); err != nil {
		s.E500(err.Error())
		return
	}

	system := &models.System{
		Name:                svo.Name,
		AppKey:              s.generateAppKey(),
		Secret:              s.generateAppKey(),
		KeepSign:            svo.KeepSign,
		Strategy:            svo.Strategy,
		TemplateRoleVisible: svo.TemplateRoleVisible,
		Status:              models.Active,
	}
	role := &models.Role{
		Name:   "owner",
		Status: models.Active,
		FromId: 1,
	}

	if o, err := business.TransactionStart(); err != nil {
		s.E500("创建系统失败")
		return
	} else if _, err := o.Insert(system); err != nil {
		business.TransactionProcess(o, err)
		s.E500(err.Error())
		return
	} else if _, err := o.Insert(role); err != nil {
		business.TransactionProcess(o, err)

		s.E500("创建角色失败")
		return
	} else {
		userRole := &models.UserRole{
			RoleId: role.Id,
			UserId: userInfo.Id,
		}
		if _, err := o.Insert(userRole); err != nil {
			business.TransactionProcess(o, err)
			s.E500("创建管理员失败")
		}
		business.TransactionEnd(o, nil)
	}

	s.Return("入驻成功")
}

func (s *SystemController) generateAppKey() string {
	s.Lock()
	defer s.Unlock()
	hash := md4.New()
	has := hash.Sum([]byte(business.CastInt64ToString(time.Now().UnixNano())))
	return fmt.Sprintf("%x", has)
}
