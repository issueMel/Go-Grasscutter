package resp

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketClientAbilityInitFinishNotify(p *player.Player) *base.Packet {
	code := base.ClientAbilityInitFinishNotify

	var entityId uint32
	if len(p.ClientAbility) > 0 {
		entityId = p.ClientAbility[0].EntityId
	}

	msg := pb.ClientAbilityInitFinishNotify{
		EntityId: entityId,
		Invokes:  p.ClientAbility,
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
