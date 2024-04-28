package command

type PermissionHandler interface {
	EnablePermissionCommand() bool
	CheckPermission() bool
}
