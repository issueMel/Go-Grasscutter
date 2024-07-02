package achievement

import (
	"Go-Grasscutter/data"
	"Go-Grasscutter/db"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
			return create(uid)
		}
		log.SugaredLogger.Error(err)
		return nil
	}
	return a
}

func create(uid int) *Achievements {
	newAchievement := &Achievements{
		Uid:                    uid,
		AchievementList:        listInit(),
		FinishedAchievementNum: 0,
		TakenGoalRewardIdList:  make([]int32, 0),
	}
	newAchievement.save()
	return newAchievement
}

func listInit() map[int]*Achievement {
	m := make(map[int]*Achievement)
	for _, v := range data.GameData.AchievementDataMap {
		if v.IsUsed() {
			m[v.Id] = &Achievement{
				Status:             pb.Achievement_Status_name[int32(pb.Achievement_Status_Unfinished)],
				Id:                 v.Id,
				TotalProgress:      v.Progress,
				CurProgress:        0,
				FinishTimestampSec: 0,
			}
		}
	}
	return m
}

func (a *Achievements) save() {
	opts := options.Update().SetUpsert(true)
	_, err := db.DB.Collection(collName).UpdateOne(context.Background(), bson.D{{"uid", a.Uid}}, a, opts)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
}
