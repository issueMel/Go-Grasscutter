package tower

type TowerLevelRecord struct {
	FloorId                 int         `bson:"floorId"`
	PassedLevelMap          map[int]int `bson:"passedLevelMap"`
	FloorStarRewardProgress int         `bson:"floorStarRewardProgress"`
}
