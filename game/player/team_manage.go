package player

type TeamManager struct {
	Teams                 map[int]*TeamInfo `bson:"teams"`
	CurrentTeamIndex      int               `bson:"currentTeamIndex"`
	CurrentCharacterIndex int               `bson:"currentCharacterIndex"`
}
