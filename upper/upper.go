package upper

import (
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"strings"
	"tianwei.pro/business"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam/facade"
	"tianwei.pro/sam/models"
)

var (
	AppKeyOrSecretError = errors.New("请检查appKey或secret")
	SystemError = errors.New("权限系统错误")
)

type SamFilterImpl struct {
	TokenFacade facade.TokenFacade `inject:"tokenFacade"`
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

	var apis []*sam_agent.Router

	if sApis, err := models.LoadApiBySystemAndStatusAndType(system.Id, models.Active, models.Button); err != nil {
		logs.Error("load api failed. err: %v", err)
		return nil, SystemError
	} else {
		for _, api := range sApis {
			apis = append(apis, &sam_agent.Router{
				Id: api.Id,
				Url: api.Path,
				Method: api.Method,
				Type: api.VerificationType,
			})
		}
	}
	return &sam_agent.SystemInfo{
		Id: system.Id,
		PermissionType: system.Strategy,
		KeepSign: system.KeepSign,
		Routers: apis,
	}, nil
}

func (s *SamFilterImpl) VerifyToken(appKey, secret, token string) (*sam_agent.UserInfo, error) {
	system, err := s.LoadSystemInfo(appKey, secret)
	if err != nil {
		return nil, err
	}
	if user, err := facade.T.DecodeToken(token); err != nil {
		return nil, err
	} else {
		var roles []*models.Role
		userRole := &models.UserRole{ UserId: user.Id, SystemId: system.Id}
		if uroles, err := userRole.LoadByUserAndSystemId(); err != nil {
			return nil, err
		} else if len(uroles) == 0 {

		} else {
			var roleIds []int64
			for _, role := range uroles {
				roleIds = append(roleIds, role.RoleId)
			}
			if rr, err := models.LoadByRoleIdsAndSystemIdAndStatus(roleIds, system.Id, models.Active); err != nil {
				return nil, err
			} else {
				roles = rr
			}
		}

		var permissions []*sam_agent.Permission

		for _, role := range roles {
			ids := strings.Split(role.PermissionSet, ",")
			var pids []int64
			for _, id := range ids {
				pids = append(pids, business.CastStringToInt64DefaultValue(id, -1))
			}

			// todo:: branchIds
			permissions = append(permissions, &sam_agent.Permission{
				RoleId: role.Id,
				RoleName: role.Name,
				PermissionSet: pids,
			})
		}

		return &sam_agent.UserInfo{
			Id: user.Id,
			UserName: user.UserName,
			Permissions: permissions,
		}, nil
	}


}
