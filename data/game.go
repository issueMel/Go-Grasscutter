package data

import "Go-Grasscutter/data/excels/avatar"

var GameData *Data = &Data{
	AvatarDataMap: make(map[int]*avatar.Data),
}

type Data struct {
	AvatarDataMap           map[int]*avatar.Data
	AvatarSkillDepotDataMap map[int]any // todo AvatarSkillDepotData
}
