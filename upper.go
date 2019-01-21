package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam/models"
)

var (
	AppKeyOrSecretError = errors.New("请检查appKey或secret")
)

type SamFilterImpl struct {

}

func init()  {
	sam_agent.SamAgent = &SamFilterImpl{}
}

func (s *SamFilterImpl) LoadSystemInfo(appKey, secret string) (*sam_agent.SystemInfo, error) {
	system := &models.System{
		AppKey: appKey,
	}
	if err := system.FindByAppKey(); err != nil {
		logs.Warn("app key: %s not found", appKey)
		return nil, AppKeyOrSecretError
	}
	if system.Secret != secret {
		logs.Warn("app key: %s, secret: %s, sys: %v", appKey, secret, system)
		return nil, AppKeyOrSecretError
	}
	return &sam_agent.SystemInfo{
		PermissionType: system.Strategy,

	}, nil
}

func (s *SamFilterImpl) VerifyToken(token string) (*sam_agent.UserInfo, error) {
	panic("implement me")
}
