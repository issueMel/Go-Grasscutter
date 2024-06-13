package resp

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketGetPlayerSocialDetailRsp(detail *pb.SocialDetail) *base.Packet {
	code := base.GetPlayerSocialDetailRsp
	msg := pb.GetPlayerSocialDetailRsp{}

	if detail != nil {
		msg.DetailData = detail
	} else {
		msg.Retcode = int32(pb.Retcode_RET_SVR_ERROR)
	}

	data, err := proto.Marshal(&msg)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	return &base.Packet{
		Opcode: code,
		Data:   data,
	}
}
