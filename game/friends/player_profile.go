package friends

// todo CHECK: all can get from player, need PlayerProfile ?
type PlayerProfile struct {
	Uid              int    `bson:"uid"`
	NameCard         int    `bson:"nameCard"`
	AvatarId         int    `bson:"avatarId"`
	Name             string `bson:"name"`
	Signature        string `bson:"signature"`
	PlayerLevel      int    `bson:"playerLevel"`
	WorldLevel       int    `bson:"worldLevel"`
	LastActiveTime   int    `bson:"lastActiveTime"`
	IsInDuel         bool   `bson:"isInDuel"`
	IsDuelObservable bool   `bson:"isDuelObservable"`
	EnterHomeOption  int    `bson:"enterHomeOption"` // todo INCOMPLETE: GameHome
}
