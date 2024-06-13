package resp

import (
	"Go-Grasscutter/game/quest"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketGetAllActivatedBargainDataRsp(records map[int]*quest.BargainRecord) *base.Packet {
	code := base.GetAllActivatedBargainDataRsp
	l := len(records)
	msg := pb.GetAllActivatedBargainDataRsp{
		SnapshotList: make([]*pb.BargainSnapshot, l),
		Retcode:      int32(pb.Retcode_RET_SUCC),
	}

	idx := 0
	for _, value := range records {
		msg.SnapshotList[idx] = value.ToSnapshot()
		idx++
	}

	data, err := proto.Marshal(&msg)
	if err != nil {
		log.SugaredLogger.Error(err)
	}

	return &base.Packet{
		Opcode: code,
		Data:   data,
	}
}
