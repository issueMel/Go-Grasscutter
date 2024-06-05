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
	session.Router[base.PlayerSetPauseReq] = HandlerPlayerSetPauseReq
}

func HandlerPlayerSetPauseReq(sess *session.Session, header, payload []byte) {
	req := pb.PlayerSetPauseReq{}
	err := proto.Unmarshal(payload, &req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}

	// Check if the player is in a multiplayer world.
	// todo INCOMPLETE
	// sess.Send(resp.PacketPlayerSetPauseRsp(pb.Retcode_RET_FAIL))
	sess.Send(resp.PacketPlayerSetPauseRsp(pb.Retcode_RET_SUCC))
}
