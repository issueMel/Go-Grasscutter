package resp

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/data"
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/proto"
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
		RewardTakenGoalIdList: make([]uint32, 0),
		AchievementList:       make([]*pb.Achievement, 0),
	}

	err := copier.Copy(&msg.RewardTakenGoalIdList, achievements.TakenGoalRewardIdList)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
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
