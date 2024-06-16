package resp

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketActivityScheduleInfoNotify() *base.Packet {
	code := base.ActivityScheduleInfoNotify
	msg := pb.ActivityScheduleInfoNotify{
		ActivityScheduleList: make([]*pb.ActivityScheduleInfo, 0),
	}
	// todo ActivityManager
	data, err := proto.Marshal(&msg)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	return &base.Packet{
		Opcode:            code,
		ShouldBuildHeader: true,
		Data:              data,
	}
}
