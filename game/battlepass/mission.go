package battlepass

type Mission struct {
	Id       int    `bson:"id"`
	Progress int    `bson:"progress"`
	Status   string `bson:"status"` // todo BattlePassMissionStatus
}
