package system

type RouterGroup struct {
	BaseRouter
	UserRouter
	RoleRouter
	PermissionRouter
	OperationLogRouter
}
