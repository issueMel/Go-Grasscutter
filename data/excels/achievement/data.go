package achievement

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/data/excels"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
	"sync/atomic"
)

const achievementResources = "AchievementExcelConfigData.json"

type Data struct {
	isDivided                  atomic.Bool
	groupAchievementIdList     map[int]struct{}
	isParent                   bool
	GoalId                     int                   `json:"goalId,omitempty"`
	PreStageAchievementId      int                   `json:"preStageAchievementId,omitempty"`
	TitleTextMapHash           int64                 `json:"titleTextMapHash,omitempty"`
	DescTextMapHash            int64                 `json:"descTextMapHash,omitempty"`
	FinishRewardId             int                   `json:"finishRewardId,omitempty"`
	IsDeleteWatcherAfterFinish bool                  `json:"isDeleteWatcherAfterFinish,omitempty"`
	Id                         int                   `json:"id,omitempty"`
	TriggerConfig              *excels.TriggerConfig `json:"triggerConfig,omitempty"`
	Progress                   int                   `json:"progress,omitempty"`
	IsDisuse                   bool                  `json:"isDisuse,omitempty"`
}

func LoadAchievementData() map[int]*Data {
	data := []*Data{}
	prefix := config.Conf.FolderStructure.Resources
	fp := filepath.Join(prefix + "ExcelBinOutput/" + achievementResources)
	file, err := os.ReadFile(fp)
	if err != nil {
		log.SugaredLogger.Panic(err)
		return nil
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.SugaredLogger.Panic(err)
		return nil
	}

	dataMap := make(map[int]*Data)

	for _, v := range data {
		dataMap[v.Id] = v
	}
	return dataMap
}

func (data *Data) IsUsed() bool {
	return !data.IsDisuse
}

func (data *Data) divideIntoGroups() {
	// todo
}
