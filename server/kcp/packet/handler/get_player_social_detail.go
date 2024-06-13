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
	session.Router[base.GetPlayerSocialDetailReq] = HandlerGetPlayerSocialDetailReq
}

func HandlerGetPlayerSocialDetailReq(sess *session.Session, header, payload []byte) {
	req := pb.GetPlayerTokenReq{}
	err := proto.Unmarshal(payload, &req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
	// todo getSocialDetailByUid()
	detail := sess.Player.GetSocialDetail()
	if detail != nil {
		detail.IsFriend = true // todo FriendsList
	}
	sess.Send(resp.PacketGetPlayerSocialDetailRsp(detail))
}
