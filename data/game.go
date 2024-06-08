package data

import (
	"Go-Grasscutter/data/excels"
	"Go-Grasscutter/data/excels/avatar"
	"Go-Grasscutter/data/excels/scene"
)

var GameData *Data = &Data{
	AvatarDataMap:            avatar.LoadAvatarData(),
	AvatarSkillDepotDataMap:  avatar.LoadAvatarSkillDepotData(),
	BattlePassMissionDataMap: excels.LoadBattlePassMissionData(),
	SceneTagDataMap:          scene.LoadSceneTagData(),
	OpenStateList:            excels.LoadOpenState(),
	ItemDataMap:              excels.LoadItemData(),
}

type Data struct {
	AvatarDataMap            map[int]*avatar.Data
	AvatarSkillDepotDataMap  map[int]*avatar.SkillDepotData
	BattlePassMissionDataMap map[int]*excels.BattlePassMissionData
	SceneTagDataMap          map[int]*scene.TagData
	OpenStateList            []*excels.OpenStateData
	ItemDataMap              map[int]*excels.ItemData
}
