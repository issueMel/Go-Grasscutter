package enum

type ParentQuestState uint32

const (
	PARENT_QUEST_STATE_NONE     ParentQuestState = 0
	PARENT_QUEST_STATE_FINISHED ParentQuestState = 1
	PARENT_QUEST_STATE_FAILED   ParentQuestState = 2
	PARENT_QUEST_STATE_CANCELED ParentQuestState = 3
)

// ParentQuestState_name if value dont exist, will return ""
var ParentQuestState_name = map[ParentQuestState]string{
	0: "PARENT_QUEST_STATE_NONE",
	1: "PARENT_QUEST_STATE_FINISHED",
	2: "PARENT_QUEST_STATE_FAILED",
	3: "PARENT_QUEST_STATE_CANCELED",
}

// ParentQuestState_value if value dont exist, will return 0
var ParentQuestState_value = map[string]uint32{
	"PARENT_QUEST_STATE_NONE":     0,
	"PARENT_QUEST_STATE_FINISHED": 1,
	"PARENT_QUEST_STATE_FAILED":   2,
	"PARENT_QUEST_STATE_CANCELED": 3,
}
