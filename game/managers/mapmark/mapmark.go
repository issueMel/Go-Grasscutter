package mapmark

import (
	"Go-Grasscutter/game/world"
	"Go-Grasscutter/generated/pb"
)

type MapMark struct {
	SceneId          int                  `bson:"sceneId"`
	Name             string               `bson:"name"`
	Position         *world.Position      `bson:"position"`
	MapMarkPointType *pb.MapMarkPointType `bson:"mapMarkPointType"`
	MonsterId        int                  `bson:"monsterId"`
	MapMarkFromType  *pb.MapMarkFromType  `bson:"mapMarkFromType"`
	QuestId          int                  `bson:"questId"`
}
