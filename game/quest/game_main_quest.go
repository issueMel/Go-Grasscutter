package quest

import (
	"Go-Grasscutter/data/binout"
	"Go-Grasscutter/generated/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameMainQuest struct {
	Id                        primitive.ObjectID       `bson:"id"`
	OwnerUid                  int                      `bson:"ownerUid"`
	ChildQuests               map[int]*GameQuest       `bson:"childQuests"`
	ParentQuestId             int                      `bson:"parentQuestId"`
	QuestVars                 []int                    `bson:"questVars"`
	TimeVar                   []int64                  `bson:"timeVar"`
	State                     string                   `bson:"state"` // enum.ParentQuestState
	IsFinished                bool                     `bson:"isFinished"`
	QuestGroupSuites          []*GroupSuite            `bson:"questGroupSuites"`
	SuggestTrackMainQuestList []int                    `bson:"suggestTrackMainQuestList"`
	Talks                     map[int]*binout.TalkData `bson:"talks"`
}

func (g *GameMainQuest) ToProto(withChildQuests bool) *pb.ParentQuest {
	// todo INCOMPLETE: ParentQuest.ToProto()
	return &pb.ParentQuest{}
}
