package resp

import (
	"Go-Grasscutter/data"
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketGetScenePointRsp(p *player.Player, sceneId uint32) *base.Packet {
	code := base.GetScenePointRsp
	msg := pb.GetScenePointRsp{
		UnlockedPointList: make([]uint32, 0),
		UnhidePointList:   make([]uint32, 0),
		UnlockAreaList:    make([]uint32, 0),
	}

	if len(data.GameData.ScenePointIdList) == 0 {
		for i := 0; i < 1000; i++ {
			msg.UnlockedPointList = append(msg.UnlockedPointList, uint32(i))
			msg.UnhidePointList = append(msg.UnhidePointList, uint32(i))
		}
	} else {
		// todo data.GameData.ScenePointIdList > 0
	}

	for i := 1; i < 9; i++ {
		msg.UnlockAreaList = append(msg.UnlockAreaList, uint32(i))
	}

	d, err := proto.Marshal(&msg)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	return &base.Packet{
		Opcode: code,
		Data:   d,
	}
}
