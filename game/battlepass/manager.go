package battlepass

import "go.mongodb.org/mongo-driver/bson/primitive"

type Manager struct {
	Id           primitive.ObjectID `bson:"id"`
	OwnerUid     int                `bson:"ownerUid"`
	Point        int                `bson:"point"`
	CyclePoints  int                `bson:"cyclePoints"`
	Level        int                `bson:"level"`
	Viewed       bool               `bson:"viewed"`
	Paid         bool               `bson:"paid"`
	Missions     map[int]*Mission   `bson:"missions"`
	TakenRewards map[int]*Reward    `bson:"takenRewards"`
}
