package player

import "Go-Grasscutter/game/quest"

type PlayerProgress struct {
	ItemHistory map[int]*ItemEntry `bson:"itemHistory"`

	// A list of dungeon IDs which have been completed.
	// This only applies to one-time dungeons.
	CompletedDungeons []int `bson:"completedDungeons"`

	// keep track of EXEC_ADD_QUEST_PROGRESS count, will be used in CONTENT_ADD_QUEST_PROGRESS
	// not sure where to put this, this should be saved to DB but not to individual quest, since
	// it will be hard to loop and compare
	QuestProgressCountMap map[string]int `bson:"questProgressCountMap"`

	ItemGivings map[int]*quest.ItemGiveRecord `bson:"itemGivings"`
	Bargains    map[int]*quest.BargainRecord  `bson:"bargains"`
}

type ItemEntry struct {
	ItemId        int `bson:"itemId"`
	ObtainedCount int `bson:"obtainedCount"`
}
