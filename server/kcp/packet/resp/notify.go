package resp

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
	"time"
)

func PacketPlayerDataNotify(player *player.Player) *base.Packet {
	code := base.PlayerDataNotify

	msg := pb.PlayerDataNotify{
		NickName:          player.Nickname,
		ServerTime:        uint64(time.Now().UnixMilli()),
		IsFirstLoginToday: true,
		RegionId:          uint32(player.RegionID),
		PropMap:           make(map[uint32]*pb.PropValue),
	}

	for key, val := range player.Properties {
		msg.PropMap[uint32(key)] = &pb.PropValue{
			Type: uint32(key),
			Value: &pb.PropValue_Ival{
				Ival: int64(val),
			},
			Val: int64(val),
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
	packet.BuildHeader(2)
	return packet
}

func PacketStoreWeightLimitNotify() *base.Packet {
	code := base.StoreWeightLimitNotify
	// todo need config ?
	msg := pb.StoreWeightLimitNotify{
		WeightLimit:         30000,
		FurnitureCountLimit: 2000,
		WeaponCountLimit:    20000,
		ReliquaryCountLimit: 2000,
		MaterialCountLimit:  2000,
		StoreType:           pb.StoreType_STORE_TYPE_PACK,
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

func PacketPlayerStoreNotify(player *player.Player) *base.Packet {
	code := base.PlayerStoreNotify

	msg := pb.PlayerStoreNotify{
		StoreType:   pb.StoreType_STORE_TYPE_PACK,
		WeightLimit: 30000,
		ItemList:    make([]*pb.Item, 0),
	}

	for _, item := range player.Inventory.Store {
		msg.ItemList = append(msg.ItemList, item.ToProto())
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
	packet.BuildHeader(2)
	return packet
}

func PacketAvatarDataNotify(player *player.Player) *base.Packet {
	code := base.AvatarDataNotify

	msg := pb.AvatarDataNotify{
		CurAvatarTeamId:   uint32(player.TeamManager.CurrentTeamIndex),
		ChooseAvatarGuid:  uint64(player.TeamManager.CurrentCharacterIndex),
		OwnedFlycloakList: []uint32{},
		OwnedCostumeList:  []uint32{},
		AvatarList:        make([]*pb.AvatarInfo, 0),
		AvatarTeamMap:     make(map[uint32]*pb.AvatarTeam),
	}
	err := copier.Copy(&msg.OwnedFlycloakList, player.FlyCloakList)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}
	err = copier.Copy(msg.OwnedCostumeList, player.CostumeList)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}
	for _, val := range player.Avatars.Avatars {
		msg.AvatarList = append(msg.AvatarList, val.ToProto())
	}

	for key, val := range player.TeamManager.Teams {
		// todo change val
		msg.AvatarTeamMap[uint32(key)] = val.ToProto()
	}
	// Set main character
	val, ok := player.Avatars.Avatars[player.MainCharacterID]
	if ok {
		msg.ChooseAvatarGuid = uint64(val.Guid)
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
