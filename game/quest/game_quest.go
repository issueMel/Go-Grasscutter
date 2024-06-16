package quest

import (
	"Go-Grasscutter/data/excels/quest"
	"Go-Grasscutter/game/quest/enum"
	"Go-Grasscutter/generated/pb"
	"unsafe"
)

type GameQuest struct {
	// GameMainQuest
	QuestData          *quest.Data
	SubQuestId         int    `bson:"subQuestId"`
	MainQuestId        int    `bson:"mainQuestId"`
	State              string `bson:"state"` // enum.QuestState
	StartTime          int    `bson:"startTime"`
	AcceptTime         int    `bson:"acceptTime"`
	FinishTime         int    `bson:"finishTime"`
	StartGameDay       int64  `bson:"startGameDay"`
	FinishProgressList []int  `bson:"finishProgressList"`
	FailProgressList   []int  `bson:"failProgressList"`
	// triggerData
	Triggers map[string]bool `bson:"triggers"`
}

func (g *GameQuest) ToProto() *pb.Quest {
	prot := &pb.Quest{
		QuestId:            uint32(g.SubQuestId),
		State:              enum.ParentQuestState_value[g.State],
		ParentQuestId:      uint32(g.MainQuestId),
		StartTime:          uint32(g.StartTime),
		StartGameTime:      438,
		AcceptTime:         uint32(g.AcceptTime),
		FinishProgressList: *(*[]uint32)(unsafe.Pointer(&g.FinishProgressList)),
		FailProgressList:   *(*[]uint32)(unsafe.Pointer(&g.FailProgressList)),
	}
	return prot
}
