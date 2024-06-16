package handler

import (
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/session"
)

func init() {
	session.Router[base.GetPlayerBlacklistReq] = HandlerGetPlayerBlacklistReq
}

func HandlerGetPlayerBlacklistReq(sess *session.Session, header, payload []byte) {
	packet := base.NewPacketWithCode(base.GetPlayerBlacklistRsp)
	packet.BuildHeader(3)
	sess.Send(packet)
}
