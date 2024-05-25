package resp

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketPingRsp(clientSeq, time uint32) *base.Packet {
	code := base.PingRsp

	p := pb.PingRsp{
		ClientTime: time,
	}
	data, err := proto.Marshal(&p)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	packet := &base.Packet{
		Opcode: code,
		Data:   data,
	}
	packet.BuildHeader(clientSeq)

	return packet
}
