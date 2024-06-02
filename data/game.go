package data

import "Go-Grasscutter/data/excels/avatar"

var GameData *Data = &Data{
	AvatarDataMap: make(map[int]*avatar.AvatarData),
}

type Data struct {
	AvatarDataMap           map[int]*avatar.AvatarData
	AvatarSkillDepotDataMap map[int]any // todo AvatarSkillDepotData
}
