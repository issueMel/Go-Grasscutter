package quest

type GameQuest struct {
	// GameMainQuest
	// QuestData
	SubQuestId         int    `bson:"subQuestId"`
	MainQuestId        int    `bson:"mainQuestId"`
	State              string `bson:"state"` // todo QuestState
	StartTime          int    `bson:"startTime"`
	AcceptTime         int    `bson:"acceptTime"`
	FinishTime         int    `bson:"finishTime"`
	StartGameDay       int64  `bson:"startGameDay"`
	FinishProgressList []int  `bson:"finishProgressList"`
	FailProgressList   []int  `bson:"failProgressList"`
	// triggerData
	Triggers map[string]bool `bson:"triggers"`
}
