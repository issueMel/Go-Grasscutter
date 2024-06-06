package enum

type QuestState uint32

const (
	QUEST_STATE_NONE       QuestState = 0
	QUEST_STATE_UNSTARTED  QuestState = 1
	QUEST_STATE_UNFINISHED QuestState = 2
	QUEST_STATE_FINISHED   QuestState = 3
	QUEST_STATE_FAILED     QuestState = 4
)

// Used by lua
const (
	NONE = iota
	UNSTARTED
	UNFINISHED
	FINISHED
	FAILED
)

var QuestState_name = map[QuestState]string{
	0: "QUEST_STATE_NONE",
	1: "QUEST_STATE_UNSTARTED",
	2: "QUEST_STATE_UNFINISHED",
	3: "QUEST_STATE_FINISHED",
	4: "QUEST_STATE_FAILED",
}

var QuestState_value = map[string]uint32{
	"QUEST_STATE_NONE":       0,
	"QUEST_STATE_UNSTARTED":  1,
	"QUEST_STATE_UNFINISHED": 2,
	"QUEST_STATE_FINISHED":   3,
	"QUEST_STATE_FAILED":     4,

	"NONE":       0,
	"UNSTARTED":  1,
	"UNFINISHED": 2,
	"FINISHED":   3,
	"FAILED":     4,
}
