package upper

import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/server"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam/facade"
	"tianwei.pro/sam/models"
)

var (
	AppKeyOrSecretError = errors.New("请检查appKey或secret")
	SystemError         = errors.New("权限系统错误")
)

var  SamAgentFacade *SamAgentFacadeImpl

type SamAgentFacadeImpl struct {
}

func init()  {
	s := server.NewServer()
	addRegistryPlugin(s)

	SamAgentFacade = new(SamAgentFacadeImpl)
	s.RegisterName("SamAgentFacadeImpl", SamAgentFacade, "")

	go func() {
		s.Serve("tcp", "localhost:0")
	}()
}

func addRegistryPlugin(s *server.Server) {
	r := client.InprocessClient
	s.Plugins.Add(r)
}


func (s *SamAgentFacadeImpl) verifySecret(appKey, secret string) (*models.System, error) {
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

func (s *SamAgentFacadeImpl) LoadSystemInfo(ctx context.Context, param *sam_agent.SystemInfoParam, reply *sam_agent.SystemInfo) error {
	system, err := s.verifySecret(param.AppKey, param.Secret)
	if err != nil {
		return err
	}

	var apis []*sam_agent.Router

	if sApis, err := models.LoadApiBySystemAndStatusAndType(system.Id, models.Active, models.Button); err != nil {
		logs.Error("load api failed. err: %v", err)
		return SystemError
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

	reply.Id = system.Id
	reply.PermissionType = system.Strategy
	reply.KeepSign = system.KeepSign
	reply.Routers = apis
	return nil
}

func (s *SamAgentFacadeImpl) VerifyToken(ctx context.Context, param *sam_agent.VerifyTokenParam, reply *sam_agent.UserInfo) error {
	system, err := s.verifySecret(param.AppKey, param.Secret)
	if err != nil {
		return err
	}
	if user, err := facade.T.DecodeToken(param.Token); err != nil {
		return err
	} else if r, err := facade.LoadAgentUserInfo(user.Id, system, user.UserName); err != nil {
		return err
	} else {
		reply.Id = r.Id
		reply.UserName = r.UserName
		reply.Avatar = r.Avatar
		reply.Email = r.Email
		reply.Phone = r.Phone
		reply.Permissions = r.Permissions
		return nil
	}
}
