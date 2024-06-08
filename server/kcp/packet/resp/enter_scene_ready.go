package resp

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
)

func PacketEnterScenePeerNotify(player *player.Player) *base.Packet {
	code := base.EnterScenePeerNotify

	msg := pb.EnterScenePeerNotify{
		DestSceneId:     uint32(player.SceneID),
		PeerId:          uint32(player.PeerId),
		HostPeerId:      uint32(player.PeerId), // todo CHECK
		EnterSceneToken: uint32(player.EnterSceneToken),
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

func PacketEnterSceneReadyRsp(player *player.Player) *base.Packet {
	code := base.EnterSceneReadyRsp

	msg := pb.EnterSceneReadyRsp{
		EnterSceneToken: uint32(player.EnterSceneToken),
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
	packet.BuildHeader(11)

	return packet
}
