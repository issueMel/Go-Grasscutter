package data

import (
	"Go-Grasscutter/data/excels"
	"Go-Grasscutter/data/excels/avatar"
)

var GameData *Data = &Data{
	AvatarDataMap:            make(map[int]*avatar.Data),
	AvatarSkillDepotDataMap:  make(map[int]*avatar.SkillDepotData),
	BattlePassMissionDataMap: make(map[int]*excels.BattlePassMissionData),
}

type Data struct {
	AvatarDataMap            map[int]*avatar.Data
	AvatarSkillDepotDataMap  map[int]*avatar.SkillDepotData
	BattlePassMissionDataMap map[int]*excels.BattlePassMissionData
}
