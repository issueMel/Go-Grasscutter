package resp

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketBattlePassMissionUpdateNotify(p *player.Player) *base.Packet {
	code := base.BattlePassMissionUpdateNotify
	msg := pb.BattlePassMissionUpdateNotify{}
	if p.BattlePassManager.Missions != nil {
		msg.MissionList = make([]*pb.BattlePassMission, len(p.BattlePassManager.Missions))
		idx := 0
		for _, val := range p.BattlePassManager.Missions {
			msg.MissionList[idx] = val.ToProto()
			idx++
		}
	}

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
