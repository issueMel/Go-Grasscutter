package quest

import "Go-Grasscutter/data"

type Manager struct {
	Uid                 int
	MainQuests          map[int]*GameMainQuest
	AcceptProgressLists map[int][]int
	LoggedQuests        []int
	LastHourCheck       int64
	LastDayCheck        int64
}

func (m *Manager) LoadFromDatabase() {
	quests := GetAllQuests(m.Uid)

	for _, mainQuest := range quests {
		cancelAdd := false
		for _, quest := range mainQuest.ChildQuests {
			questConfig, ok := data.GameData.QuestDataMap[quest.SubQuestId]
			if !ok {
				// todo remove form db
				cancelAdd = true
				break
			}

			quest.QuestData = questConfig
		}

		if !cancelAdd {
			m.MainQuests[mainQuest.ParentQuestId] = mainQuest
		}
	}
}
