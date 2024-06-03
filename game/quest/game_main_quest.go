package quest

import (
	"Go-Grasscutter/data/binout"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameMainQuest struct {
	Id                        primitive.ObjectID       `bson:"id"`
	OwnerUid                  int                      `bson:"ownerUid"`
	ChildQuests               map[int]*GameQuest       `bson:"childQuests"`
	ParentQuestId             int                      `bson:"parentQuestId"`
	QuestVars                 []int                    `bson:"questVars"`
	TimeVar                   []int64                  `bson:"timeVar"`
	State                     string                   `bson:"state"` // todo ParentQuestState
	IsFinished                bool                     `bson:"isFinished"`
	QuestGroupSuites          []*GroupSuite            `bson:"questGroupSuites"`
	SuggestTrackMainQuestList []int                    `bson:"suggestTrackMainQuestList"`
	Talks                     map[int]*binout.TalkData `bson:"talks"`
}
