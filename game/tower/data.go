package tower

type TowerData struct {
	currentFloorId int                      `bson:"currentFloorId"`
	currentLevel   int                      `bson:"currentLevel"`
	recordMap      map[int]TowerLevelRecord `bson:"recordMap"`
}
