package handler

import (
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/packet/resp"
	"Go-Grasscutter/server/kcp/session"
)

func init() {
	session.Router[base.EnterSceneReadyReq] = HandlerEnterSceneReadyReq
}

func HandlerEnterSceneReadyReq(sess *session.Session, header, payload []byte) {
	sess.Send(resp.PacketEnterScenePeerNotify(sess.Player))
	sess.Send(resp.PacketEnterSceneReadyRsp(sess.Player))
}
