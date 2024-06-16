package handler

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/packet/resp"
	"Go-Grasscutter/server/kcp/session"
)

func init() {
	session.Router[base.SceneInitFinishReq] = HandlerSceneInitFinishReq
}

func HandlerSceneInitFinishReq(sess *session.Session, header, payload []byte) {
	p := sess.Player

	// todo Info packets
	sess.Send(resp.PacketServerTimeNotify())
	sess.Send(resp.PacketWorldPlayerInfoNotify())
	sess.Send(resp.PacketWorldDataNotify(p))
	sess.Send(resp.PacketPlayerWorldSceneInfoListNotify(p))
	sess.Send(base.NewPacketWithCode(base.SceneForceUnlockNotify))

	sess.Send(resp.PacketSceneTimeNotify(p))
	sess.Send(resp.PacketPlayerGameTimeNotify(p))
	sess.Send(resp.PacketPlayerEnterSceneInfoNotify(p))
	sess.Send(resp.PacketSceneAreaWeatherNotify(p))
	sess.Send(resp.PacketScenePlayerInfoNotify(p))
	sess.Send(resp.PacketSceneTeamUpdateNotify(p))

	sess.Send(resp.PacketSyncTeamEntityNotify(p))
	sess.Send(resp.PacketSyncScenePlayTeamEntityNotify(p))

	// Done Packet
	sess.Send(resp.PacketSceneInitFinishRsp(p))

	// Set scene load state.
	p.SceneLoadState = player.INIT
	// todo Run scene initialization.

}
