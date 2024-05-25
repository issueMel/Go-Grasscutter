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
	session.Router[base.PingReq] = HandlerPingReq
}

func HandlerPingReq(sess *session.Session, header, payload []byte) {
	head := &pb.PacketHead{}
	err := proto.Unmarshal(header, head)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}

	ping := &pb.PingReq{}
	err = proto.Unmarshal(payload, ping)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}

	sess.UpdateLastPingTime(ping.GetClientTime())

	sess.Send(resp.PacketPingRsp(head.GetClientSequenceId(), ping.GetClientTime()))
}
