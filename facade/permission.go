package facade

import (
	"strings"
	"tianwei.pro/business"
	"tianwei.pro/sam-agent"
	"tianwei.pro/sam/facade/role"
	"tianwei.pro/sam/models"
)

func LoadAgentUserInfo(uid int64, system *models.System, uName string) (*sam_agent.UserInfo, error) {

	uRoles, err := role.LoadRoleIdsByUId(uid, system.Id)
	if err != nil {
		return nil, err
	}

	roles, err := findOwnRoles(uRoles, system.Id)
	if err != nil {
		return nil, err
	}

	dataPermissionMap := make(map[int64][]int64)
	if uroles, err := role.LoadRolesByUIdAndSID(uid, system.Id); err != nil {
		return nil, err
	} else {
		for _, ur := range uroles {
			data := strings.Split(ur.BranchIds, ",")
			var branchIds []int64
			for _, d := range data {
				branchIds = append(branchIds, business.CastStringToInt64(d))
			}
			dataPermissionMap[ur.RoleId] = branchIds
		}
	}

	var permissions []*sam_agent.Permission

	for _, ur := range roles {
		// 拥有的资源idx列表.
		ids := strings.Split(ur.PermissionSet, ",")
		var pids []int64
		for _, id := range ids {
			pids = append(pids, business.CastStringToInt64DefaultValue(id, -1))
		}

		var branchIds []int64

		if system.Strategy == models.DataPermissionAndOperationAuthorityLevel {
			branchIds = dataPermissionMap[ur.Id]
		} else if system.Strategy == models.TreeDataOperationPermission {
			branchIds = []int64{ur.BranchId}
		}

		for _, branchId := range branchIds {
			if childrens, err := FindByPid(branchId); err != nil {
				return nil, err
			} else {
				branchIds = append(branchIds, childrens...)
			}
		}

		permissions = append(permissions, &sam_agent.Permission{
			RoleId:        ur.Id,
			RoleName:      ur.Name,
			PermissionSet: pids,
			BranchIds:     branchIds,
		})
	}

	return &sam_agent.UserInfo{
		Id:          uid,
		UserName:    uName,
		Permissions: permissions,
	}, nil
}

func FindOwnRoles(uid, sid int64) ([]*models.Role, error){
	uRoles, err := role.LoadRoleIdsByUId(uid, sid)
	if err != nil {
		return nil, err
	}

	roles, err := findOwnRoles(uRoles, sid)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func findOwnRoles(roleIds []int64, sid int64) ([]*models.Role, error) {
	if allRoles, err := role.LoadRolesBySystemId(sid); err != nil {
		return nil, err
	} else {
		roleMap := make(map[int64]*models.Role)

		for _, role := range allRoles {
			roleMap[role.Id] = role
		}

		var ownRoles []*models.Role

		for _, roleId := range roleIds {
			if role, exist := roleMap[roleId]; exist {
				ownRoles = append(ownRoles, role)
			}
		}
		return ownRoles, nil
	}

}
