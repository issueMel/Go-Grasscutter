package handler

import (
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/packet/resp"
	"Go-Grasscutter/server/kcp/session"
)

func init() {
	session.Router[base.GetAllActivatedBargainDataReq] = HandlerGetAllActivatedBargainDataReq
}

func HandlerGetAllActivatedBargainDataReq(sess *session.Session, header, payload []byte) {
	sess.Send(resp.PacketGetAllActivatedBargainDataRsp(sess.Player.PlayerProgress.Bargains))
}
