package quest

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/data/common"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
)

const questResources = "QuestExcelConfigData.json"

type Data struct {
	SubId           int   `json:"subId,omitempty"`
	MainId          int   `json:"mainId,omitempty"`
	Order           int   `json:"order,omitempty"`
	DescTextMapHash int64 `json:"descTextMapHash,omitempty"`
	FinishParent    bool  `json:"finishParent,omitempty"`
	IsRewind        bool  `json:"isRewind,omitempty"`

	// todo LogicType
	AcceptCondComb string `json:"acceptCondComb,omitempty"`
	FinishCondComb string `json:"finishCondComb,omitempty"`
	FailCondComb   string `json:"failCondComb,omitempty"`

	AcceptCond      []*QuestCondition       `json:"acceptCond,omitempty"`
	FinishCond      []*QuestCondition       `json:"finishCond,omitempty"`
	FailCond        []*QuestCondition       `json:"failCond,omitempty"`
	BeginExec       []*QuestExecParam       `json:"beginExec,omitempty"`
	FinishExec      []*QuestExecParam       `json:"finishExec,omitempty"`
	FailExec        []*QuestExecParam       `json:"failExec,omitempty"`
	Guide           *Guide                  `json:"guide,omitempty"`
	TrialAvatarList []int                   `json:"trialAvatarList,omitempty"`
	GainItems       []*common.ItemParamData `json:"gainItems,omitempty"`
}

type QuestCondition struct {
	Type     string `json:"type,omitempty"` // todo enum QuestCond / QuestContent
	Param    []int  `json:"param,omitempty"`
	ParamStr string `json:"param_str,omitempty"`
	Count    int    `json:"count,omitempty"`
}

type QuestExecParam struct {
	Type  string `json:"type,omitempty"`  // todo enum QuestExec
	Param []any  `json:"param,omitempty"` // todo CHECK: []string panic
	Count string `json:"count,omitempty"`
}

type Guide struct {
	Type       string   `json:"type,omitempty"`
	Param      []string `json:"param,omitempty"`
	GuideScene int      `json:"guideScene,omitempty"`
}

func LoadQuestData() map[int]*Data {
	data := []*Data{}
	prefix := config.Conf.FolderStructure.Resources
	fp := filepath.Join(prefix + "ExcelBinOutput/" + questResources)
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

	sceneTagDataMap := make(map[int]*Data)

	for _, v := range data {
		sceneTagDataMap[v.SubId] = v
	}
	return sceneTagDataMap
}
