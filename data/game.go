package data

import (
	"Go-Grasscutter/data/excels"
	"Go-Grasscutter/data/excels/avatar"
	"Go-Grasscutter/data/excels/scene"
)

var GameData *Data = &Data{
	AvatarDataMap:            make(map[int]*avatar.Data),
	AvatarSkillDepotDataMap:  make(map[int]*avatar.SkillDepotData),
	BattlePassMissionDataMap: excels.LoadBattlePassMissionData(),
	SceneTagDataMap:          scene.LoadSceneTagData(),
	OpenStateList:            excels.LoadOpenState(),
}

type Data struct {
	AvatarDataMap            map[int]*avatar.Data
	AvatarSkillDepotDataMap  map[int]*avatar.SkillDepotData
	BattlePassMissionDataMap map[int]*excels.BattlePassMissionData
	SceneTagDataMap          map[int]*scene.TagData
	OpenStateList            []*excels.OpenStateData
}
