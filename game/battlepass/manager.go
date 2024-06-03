package battlepass

import (
	"Go-Grasscutter/generated/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

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

func (m *Manager) GetScheduleProto() *pb.BattlePassSchedule {
	currentDate := time.Now()

	var nextSundayDate time.Time
	if currentDate.Weekday() == time.Sunday {
		nextSundayDate = currentDate
	} else {
		offset := (7 - int(currentDate.Weekday())) % 7
		if offset == 0 {
			offset = 7
		}
		nextSundayDate = currentDate.AddDate(0, 0, offset)
	}

	nextSundayTime := time.Date(
		nextSundayDate.Year(),
		nextSundayDate.Month(),
		nextSundayDate.Day(),
		23,
		59,
		59,
		0,
		time.Local,
	)

	var us pb.BattlePassUnlockStatus
	if m.Paid {
		us = pb.BattlePassUnlockStatus_BATTLE_PASS_UNLOCK_STATUS_PAID
	} else {
		us = pb.BattlePassUnlockStatus_BATTLE_PASS_UNLOCK_STATUS_FREE
	}
	return &pb.BattlePassSchedule{
		ScheduleId:        2700,
		Level:             uint32(m.Level),
		Point:             uint32(m.Point),
		BeginTime:         0,
		EndTime:           2059483200,
		IsViewed:          m.Viewed,
		UnlockStatus:      us,
		PaidPlatformFlags: 2,
		CurCyclePoints:    uint32(m.CyclePoints),
		CurCycle: &pb.BattlePassCycle{
			BeginTime: 0,
			EndTime:   uint32(nextSundayTime.Unix()),
			CycleIdx:  3,
		},
	}
}
