package handler

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/packet/resp"
	"Go-Grasscutter/server/kcp/session"
	"google.golang.org/protobuf/proto"
)

func init() {
	session.Router[base.GetShopReq] = HandlerGetShopReq
}

func HandlerGetShopReq(sess *session.Session, header, payload []byte) {
	req := pb.GetShopReq{}
	err := proto.Unmarshal(payload, &req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
	sess.Send(resp.PacketGetShopRsp(sess.Player, req.ShopType))
}
