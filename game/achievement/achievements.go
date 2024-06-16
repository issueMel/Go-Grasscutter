package achievement

import (
	"Go-Grasscutter/db"
	"Go-Grasscutter/log"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "achievements"

type Achievements struct {
	Id                     primitive.ObjectID   `bson:"id"`
	Uid                    int                  `bson:"uid"`
	AchievementList        map[int]*Achievement `bson:"achievementList"`
	FinishedAchievementNum int                  `bson:"finishedAchievementNum"`
	TakenGoalRewardIdList  []int32              `bson:"takenGoalRewardIdList"`
}

func GetByPlayerUid(uid int) *Achievements {
	a := &Achievements{}
	err := db.DB.Collection(collName).FindOne(context.Background(), bson.D{{"uid", uid}}).Decode(a)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			// todo ENHANCE: creat
			return nil
		}
		log.SugaredLogger.Error(err)
		return nil
	}
	return a
}
