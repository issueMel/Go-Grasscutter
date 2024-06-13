package common

type PointData struct {
	Id                 int       `json:"id,omitempty"`
	AreaId             int       `json:"areaId,omitempty"`
	Type               string    `json:"$type,omitempty"`
	TranPos            *Position `json:"tranPos,omitempty"`
	Pos                *Position `json:"pos,omitempty"`
	Rot                *Position `json:"rot,omitempty"`
	Size               *Position `json:"size,omitempty"`
	ForbidSimpleUnlock bool      `json:"forbidSimpleUnlock,omitempty"`
	Unlocked           bool      `json:"unlocked,omitempty"`
	GroupLimit         bool      `json:"groupLimit,omitempty"`
	DungeonIds         []int     `json:"dungeonIds,omitempty"`
	DungeonRandomList  []int     `json:"dungeonRandomList,omitempty"`
	GroupIDs           []int     `json:"groupIDs,omitempty"`
	TranSceneId        int       `json:"tranSceneId,omitempty"`
}

// todo change world.Position
type Position struct {
	X float32 `bson:"x" json:"X"`
	Y float32 `bson:"y" json:"Y"`
	Z float32 `bson:"z" json:"Z"`
}

// todo UpdateDailyDungeon()
func (p *PointData) UpdateDailyDungeon() {
	if len(p.DungeonRandomList) == 0 {
		return
	}
	//newDungeons := []int{}
	//day := time.Now().Day() // todo getCurrentDayOfWeek()
	//for _, randomId := range p.DungeonRandomList {
	//	// data.gamedata.getDailyDungeonDataMap
	//
	//
	//}
}
