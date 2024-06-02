package player

type TeamManager struct {
	Teams                 map[int]*TeamInfo `bson:"teams"` // todo teaminfo
	CurrentTeamIndex      int               `bson:"currentTeamIndex"`
	CurrentCharacterIndex int               `bson:"currentCharacterIndex"`
}
