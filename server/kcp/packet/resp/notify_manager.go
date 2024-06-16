package resp

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/data"
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/game/pros"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"google.golang.org/protobuf/proto"
	"unsafe"
)

func PacketOpenStateUpdateNotify(player *player.Player) *base.Packet {
	code := base.OpenStateUpdateNotify
	msg := pb.OpenStateUpdateNotify{
		OpenStateMap: make(map[uint32]uint32),
	}

	usage := config.Conf.Server.Game.GameOptions.ResinOptions.ResinUsage

	for _, openStateData := range data.GameData.OpenStateList {
		id := openStateData.Id
		if id == 45 && !usage {
			msg.OpenStateMap[45] = 0
			continue
		}
		// If the player has an open state stored in their map,
		// then it would always override any default value
		if val, ok := player.OpenStates[id]; ok {
			msg.OpenStateMap[uint32(id)] = uint32(val)
		}
		// todo Otherwise, add the state if it is contained in the set of default open states.
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

func PacketAchievementAllDataNotify(player *player.Player) *base.Packet {
	code := base.AchievementAllDataNotify

	achievements := player.Achievements

	msg := pb.AchievementAllDataNotify{
		RewardTakenGoalIdList: *(*[]uint32)(unsafe.Pointer(&achievements.TakenGoalRewardIdList)),
		AchievementList:       make([]*pb.Achievement, 0),
	}

	for _, val := range achievements.AchievementList {
		msg.AchievementList = append(msg.AchievementList, val.ToProto())
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

func PacketForgeDataNotify() *base.Packet {
	code := base.ForgeDataNotify
	msg := pb.ForgeDataNotify{
		ForgeQueueMap: make(map[uint32]*pb.ForgeQueueData),
		MaxQueueNum:   4,
		ForgeIdList:   make([]uint32, 0),
	}

	// todo INCOMPLETE

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

func PacketResinChangeNotify(p *player.Player) *base.Packet {
	code := base.ResinChangeNotify
	msg := pb.ResinChangeNotify{
		CurBuyCount:      uint32(p.GetProperty(pros.PROP_PLAYER_RESIN)),
		NextAddTimestamp: uint32(p.NextResinRefresh),
		CurValue:         uint32(p.ResinBuyCount),
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

func PacketCookDataNotify() *base.Packet {
	code := base.CookDataNotify
	msg := pb.CookDataNotify{
		RecipeDataList: make([]*pb.CookRecipeData, 0), // todo cookingManager
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

func PacketCompoundDataNotify() *base.Packet {
	code := base.CompoundDataNotify
	msg := pb.CompoundDataNotify{
		CompoundQueueDataList: make([]*pb.CompoundQueueData, 0),
		UnlockCompoundList:    make([]uint32, 0),
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

func PacketUnlockedFurnitureFormulaDataNotify(p *player.Player) *base.Packet {
	code := base.UnlockedFurnitureFormulaDataNotify
	var msg = pb.UnlockedFurnitureFormulaDataNotify{
		IsAll:           true,
		FurnitureIdList: *(*[]uint32)(unsafe.Pointer(&p.UnlockedFurniture)),
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

func PacketUnlockedFurnitureSuiteDataNotify(p *player.Player) *base.Packet {
	code := base.UnlockedFurnitureSuiteDataNotify
	msg := pb.UnlockedFurnitureSuiteDataNotify{
		FurnitureSuiteIdList: *(*[]uint32)(unsafe.Pointer(&p.UnlockedFurnitureSuite)),
		IsAll:                true,
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
