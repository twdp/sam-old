package upper

import (
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam/facade"
	"tianwei.pro/sam/models"
)

var (
	AppKeyOrSecretError = errors.New("请检查appKey或secret")
	SystemError         = errors.New("权限系统错误")
)

type SamFilterImpl struct {
	TokenFacade facade.TokenFacade `inject:"tokenFacade"`
}

func init() {
	sam_agent.SamAgent = &SamFilterImpl{}
}

func (s *SamFilterImpl) verifySecret(appKey, secret string) (*models.System, error) {
	if system, err := facade.FindByKey(appKey); err != nil {
		logs.Warn("app key: %s not found", appKey)
		return nil, AppKeyOrSecretError
	} else {
		if system.Secret != secret {
			logs.Warn("app key: %s, secret: %s, sys: %v", appKey, secret, system)
			return nil, AppKeyOrSecretError
		}
		return system, nil
	}
}

func (s *SamFilterImpl) LoadSystemInfo(appKey, secret string) (*sam_agent.SystemInfo, error) {
	system, err := s.verifySecret(appKey, secret)
	if err != nil {
		return nil, err
	}

	var apis []*sam_agent.Router

	if sApis, err := models.LoadApiBySystemAndStatusAndType(system.Id, models.Active, models.Button); err != nil {
		logs.Error("load api failed. err: %v", err)
		return nil, SystemError
	} else {
		for _, api := range sApis {
			apis = append(apis, &sam_agent.Router{
				Id:     api.Id,
				Url:    api.Path,
				Method: api.Method,
				Type:   api.VerificationType,
			})
		}
	}
	return &sam_agent.SystemInfo{
		Id:             system.Id,
		PermissionType: system.Strategy,
		KeepSign:       system.KeepSign,
		Routers:        apis,
	}, nil
}

func (s *SamFilterImpl) VerifyToken(appKey, secret, token string) (*sam_agent.UserInfo, error) {
	system, err := s.verifySecret(appKey, secret)
	if err != nil {
		return nil, err
	}
	if user, err := facade.T.DecodeToken(token); err != nil {
		return nil, err
	} else {
		return facade.LoadAgentUserInfo(user.Id, system, user.UserName)
	}
}
