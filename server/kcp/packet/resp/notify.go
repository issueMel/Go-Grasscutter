package resp

import (
	data2 "Go-Grasscutter/data"
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

func PacketFinishedParentQuestNotify(player *player.Player) *base.Packet {
	code := base.FinishedParentQuestNotify
	msg := pb.FinishedParentQuestNotify{
		ParentQuestList: make([]*pb.ParentQuest, 0),
	}

	for _, val := range player.QuestManager.MainQuests {
		// todo ParentQuestState
		if val.State != "PARENT_QUEST_STATE_CANCELED" {
			msg.ParentQuestList = append(msg.ParentQuestList, val.ToProto(false))
		}
	}

	data, err := proto.Marshal(&msg)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	return &base.Packet{
		Opcode:            code,
		Data:              data,
		ShouldBuildHeader: true,
	}
}

func PacketBattlePassAllDataNotify(player *player.Player) *base.Packet {
	code := base.BattlePassAllDataNotify
	msg := pb.BattlePassAllDataNotify{
		HaveCurSchedule: true,
		CurSchedule:     player.BattlePassManager.GetScheduleProto(),
	}

	for _, missionData := range data2.GameData.BattlePassMissionDataMap {
		// Don`t send invalid refresh types
		if missionData.IsValidRefreshType() {
			continue
		}

		// Check if player has mission in bp manager.
		// If not, then add an empty proto from the mission data
		val, ok := player.BattlePassManager.Missions[missionData.Id]
		if ok {
			msg.MissionList = append(msg.MissionList, val.ToProto())
		} else {
			msg.MissionList = append(msg.MissionList, missionData.ToProto())
		}

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

func PacketQuestListNotify(player *player.Player) *base.Packet {
	code := base.QuestListNotify
	msg := pb.QuestListNotify{
		QuestList: make([]*pb.Quest, 0),
	}

	for _, child := range player.QuestManager.MainQuests {
		for _, val := range child.ChildQuests {
			// todo QuestState
			if val.State != "QUEST_STATE_UNSTARTED" {
				msg.QuestList = append(msg.QuestList, val.ToProto())
			}
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

func PacketQuestGlobalVarNotify(player *player.Player) *base.Packet {
	code := base.QuestGlobalVarNotify

	msg := pb.QuestGlobalVarNotify{
		VarList: make([]*pb.QuestGlobalVar, 0),
	}

	for key, val := range player.QuestGlobalVariables {
		msg.VarList = append(msg.VarList, &pb.QuestGlobalVar{
			Key:   uint32(key),
			Value: int32(val),
		})
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

func PacketCodexDataFullNotify(player *player.Player) *base.Packet {
	code := base.CodexDataFullNotify

	msg := pb.CodexDataFullNotify{
		TypeDataList: make([]*pb.CodexTypeData, 0), // todo CodexTypeData x 8
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

func PacketAllWidgetDataNotify(player *player.Player) *base.Packet {
	code := base.AllWidgetDataNotify
	msg := pb.AllWidgetDataNotify{
		LunchBoxData:                      &pb.LunchBoxData{},
		OneoffGatherPointDetectorDataList: make([]*pb.OneoffGatherPointDetectorData, 0),
		// AllOneoffGatherPointDetectorDataList
		PEOHMDJKMKO:             make([]*pb.WidgetCoolDownData, 0),
		AnchorPointList:         make([]*pb.AnchorPointData, 0),
		ClientCollectorDataList: make([]*pb.ClientCollectorData, 0),
		// AllNormalCoolDownDataList
		MCMDFPAFOIG: make([]*pb.WidgetCoolDownData, 0),
		SlotList:    make([]*pb.WidgetSlotData, 0),
	}

	if player.WidgetID != 0 {
		msg.SlotList = append(msg.SlotList,
			&pb.WidgetSlotData{
				IsActive:   true,
				MaterialId: uint32(player.WidgetID),
			},
			&pb.WidgetSlotData{
				Tag: pb.WidgetSlotTag_WIDGET_SLOT_TAG_ATTACH_AVATAR,
			})
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
