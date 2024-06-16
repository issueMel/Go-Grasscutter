package resp

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/utils"
	"google.golang.org/protobuf/proto"
)

func PacketGetSceneAreaRsp(p *player.Player, sceneId uint32) *base.Packet {
	code := base.GetSceneAreaRsp
	ints := p.UnlockedSceneAreas[int(sceneId)]

	msg := pb.GetSceneAreaRsp{
		CityInfoList: make([]*pb.CityInfo, 0, 5),
		SceneId:      sceneId,
		AreaIdList:   utils.ToUint32Slice(ints),
	}

	for i := 1; i < 6; i++ {
		// todo add nil CityInfo
		if val, ok := p.CityInfoData[i]; ok {
			msg.CityInfoList = append(msg.CityInfoList, val.ToProto())
		}
	}

	data, err := proto.Marshal(&msg)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	packet := &base.Packet{
		Opcode: code,
		Data:   data,
	}

	packet.BuildHeader(0)

	return packet
}
