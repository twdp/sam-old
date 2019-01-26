package models

// url类型
const (
	Menu = iota
	Page
	Button
)

// url需要的验证类型
const (
	Anonymous = iota
	OnlyNeedLogin
	CheckRolePermission
)

// 状态
const (
	Init = iota
	Active
	Freeze
)

// 客户端类型
const (
	PC = iota
	WXMP
)

const (
	NoDataPermission = iota
	DataPermissionAndOperationAuthorityLevel // 平级
	TreeDataOperationPermission // 树状
)