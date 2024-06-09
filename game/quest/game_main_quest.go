package quest

import (
	"Go-Grasscutter/data/binout"
	"Go-Grasscutter/db"
	"Go-Grasscutter/game/quest/enum"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collName = "quests"

type GameMainQuest struct {
	Id                        primitive.ObjectID       `bson:"id"`
	OwnerUid                  int                      `bson:"ownerUid"` // index
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
	prot := &pb.ParentQuest{
		ParentQuestId:    uint32(g.ParentQuestId),
		IsFinished:       g.IsFinished,
		ParentQuestState: enum.ParentQuestState_value[g.State],
		VideoKey:         0, // todo MainQuestEncryptionMap
		ChildQuestList:   make([]*pb.ChildQuest, 0),
		QuestVar:         make([]int32, len(g.QuestVars)),
	}

	if withChildQuests {
		for _, quest := range g.ChildQuests {
			if quest.State != enum.QuestState_name[enum.QUEST_STATE_UNSTARTED] {
				childQuest := pb.ChildQuest{
					QuestId: uint32(quest.SubQuestId),
					State:   enum.QuestState_value[quest.State],
				}
				prot.ChildQuestList = append(prot.ChildQuestList, &childQuest)
			}
		}
	}

	for i := range g.QuestVars {
		prot.QuestVar[i] = int32(g.QuestVars[i])
	}

	return prot
}

func GetAllQuests(uid int) []*GameMainQuest {
	find, err := db.DB.Collection(collName).Find(context.Background(), bson.D{{"ownerUid", uid}})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil
		}
		log.SugaredLogger.Error(err)
		return nil
	}

	list := []*GameMainQuest{}
	err = find.All(context.Background(), &list)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	return list
}
