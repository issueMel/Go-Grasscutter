package resp

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketPlayerSetPauseRsp(retCode pb.Retcode) *base.Packet {
	code := base.PlayerSetPauseRsp
	msg := pb.PlayerSetPauseRsp{
		Retcode: int32(retCode),
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
