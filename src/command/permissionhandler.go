package command

import "Go-Grasscutter/src/game/player"

type PermissionHandler interface {
	EnablePermissionCommand() bool
	CheckPermission(player, targetPlayer player.Player, permissionNode, permissionNodeTargeted string) bool
}

type DefaultPermissionHandler struct {
}

func (*DefaultPermissionHandler) EnablePermissionCommand() bool {
	return true
}

func (*DefaultPermissionHandler) CheckPermission(
	// todo fix CheckPermission
	player, targetPlayer *player.Player,
	permissionNode, permissionNodeTargeted string) bool {
	if player == nil {
		return true
	}
	account := player.Account.GetAccount()
	if player != targetPlayer {
		if len(permissionNodeTargeted) > 0 && len(account.Permissions) > 0 {
			return true
		}
	}
	return true
}
