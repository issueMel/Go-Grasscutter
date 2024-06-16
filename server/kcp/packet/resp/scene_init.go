package resp

import (
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
	"time"
)

func PacketServerTimeNotify() *base.Packet {
	code := base.ServerTimeNotify

	msg := pb.ServerTimeNotify{
		ServerTime: uint64(time.Now().UnixMilli()),
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

func PacketWorldPlayerInfoNotify() *base.Packet {
	code := base.WorldPlayerInfoNotify
	msg := pb.WorldPlayerInfoNotify{
		PlayerInfoList: make([]*pb.OnlinePlayerInfo, 0),
		PlayerUidList:  make([]uint32, 0),
	}

	// todo world getPlayer

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

func PacketWorldDataNotify(p *player.Player) *base.Packet {
	code := base.WorldDataNotify
	msg := pb.WorldDataNotify{
		WorldPropMap: make(map[uint32]*pb.PropValue),
	}

	msg.WorldPropMap[1] = &pb.PropValue{
		Type: 1,
		Value: &pb.PropValue_Ival{
			Ival: int64(p.GetWorldLevel()),
		},
		Val: int64(p.GetWorldLevel()),
	}

	isMp := 0
	if p.World.IsMultiplayer {
		isMp = 1
	}

	msg.WorldPropMap[2] = &pb.PropValue{
		Type: 2,
		Value: &pb.PropValue_Ival{
			Ival: int64(isMp),
		},
		Val: int64(isMp),
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

func PacketPlayerWorldSceneInfoListNotify(p *player.Player) *base.Packet {
	code := base.PlayerWorldSceneInfoListNotify

	msg := pb.PlayerWorldSceneInfoListNotify{
		InfoList: make([]*pb.PlayerWorldSceneInfo, 0),
	}
	msg.InfoList = append(msg.InfoList, &pb.PlayerWorldSceneInfo{
		SceneId:  1,
		IsLocked: false,
	})
	// todo data.GameData.SceneDataMap
	// todo data.GameData.MapLayerFloorDataMap
	// todo data.GameData.MapLayerGroupDataMap

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

func PacketSceneTimeNotify(p *player.Player) *base.Packet {
	// todo scene

	code := base.SceneTimeNotify
	msg := pb.SceneTimeNotify{
		SceneId:   0,
		SceneTime: 0,
		IsPaused:  false,
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

func PacketPlayerGameTimeNotify(p *player.Player) *base.Packet {
	code := base.PlayerGameTimeNotify

	msg := pb.PlayerGameTimeNotify{
		GameTime: uint32(p.PlayerGameTime), // todo use currentWorldTime
		Uid:      uint32(p.ID),
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

func PacketPlayerEnterSceneInfoNotify(p *player.Player) *base.Packet {
	code := base.PlayerEnterSceneInfoNotify
	empty := &pb.AbilitySyncStateInfo{}
	msg := pb.PlayerEnterSceneInfoNotify{
		CurAvatarEntityId: uint32(p.TeamManager.CurrentCharacterIndex), // todo getCurrentAvatarEntity().getId()
		EnterSceneToken:   uint32(p.EnterSceneToken),
		TeamEnterInfo: &pb.TeamEnterSceneInfo{
			TeamEntityId:        150995833, // todo entity EntityTeam
			AbilityControlBlock: nil,
			TeamAbilityInfo:     empty,
		},
		MpLevelEntityInfo: &pb.MPLevelEntityInfo{
			EntityId:        uint32(184550274), // todo p.World.Entity.Id
			AuthorityPeerId: uint32(p.PeerId),  // todo getHostPeerId()
			AbilityInfo:     empty,
		},
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

func PacketSceneAreaWeatherNotify(p *player.Player) *base.Packet {
	code := base.SceneAreaWeatherNotify

	msg := pb.SceneAreaWeatherNotify{
		ClimateType:   1, // todo enum ClimateType
		WeatherAreaId: uint32(p.WeatherId),
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

func PacketScenePlayerInfoNotify(p *player.Player) *base.Packet {
	code := base.ScenePlayerInfoNotify
	msg := pb.ScenePlayerInfoNotify{
		PlayerInfoList: make([]*pb.ScenePlayerInfo, 1),
	}
	// todo World

	msg.PlayerInfoList[0] = &pb.ScenePlayerInfo{
		Uid:     uint32(p.ID),
		PeerId:  uint32(p.PeerId),
		Name:    p.Nickname,
		SceneId: uint32(p.SceneID),
		OnlinePlayerInfo: &pb.OnlinePlayerInfo{
			Uid:           uint32(p.ID),
			Nickname:      p.Nickname,
			PlayerLevel:   uint32(p.GetLevel()),
			MpSettingType: pb.MpSettingType_MP_SETTING_TYPE_ENTER_AFTER_APPLY,
			NameCardId:    uint32(p.NameCardID),
			Signature:     p.Signature,
			ProfilePicture: &pb.ProfilePicture{
				AvatarId: uint32(p.HeadImage),
			},
			CurPlayerNumInWorld: 1,
		},
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

func PacketSceneTeamUpdateNotify(p *player.Player) *base.Packet {
	code := base.SceneTeamUpdateNotify
	msg := pb.SceneTeamUpdateNotify{
		IsInMp:              p.World.IsMultiplayer,
		SceneTeamAvatarList: make([]*pb.SceneTeamAvatar, 0),
	}

	// todo player.getWorld().getPlayers()
	// todo entity
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

func PacketSyncTeamEntityNotify(p *player.Player) *base.Packet {
	code := base.SyncTeamEntityNotify
	msg := pb.SyncTeamEntityNotify{
		SceneId:            uint32(p.SceneID),
		TeamEntityInfoList: make([]*pb.TeamEntityInfo, 0),
	}

	if p.World.IsMultiplayer {
		// todo World
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

func PacketSyncScenePlayTeamEntityNotify(p *player.Player) *base.Packet {
	code := base.SyncScenePlayTeamEntityNotify
	msg := pb.SyncScenePlayTeamEntityNotify{
		SceneId: uint32(p.SceneID),
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

func PacketSceneInitFinishRsp(player *player.Player) *base.Packet {
	code := base.SceneInitFinishRsp

	msg := pb.SceneInitFinishRsp{
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
