package scene

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
)

const SceneTagDataResource = "SceneTagConfigData.json"

type TagData struct {
	Id             int                  `json:"id"`
	Idk1           bool                 `json:"DJCOAOBDIHP,omitempty"`
	Idk2           bool                 `json:"LOLNNMPKHIB,omitempty"`
	IsDefaultValid bool                 `json:"isDefaultValid,omitempty"`
	SceneTagName   string               `json:"sceneTagName"`
	SceneId        int                  `json:"sceneId"`
	Cond           []*SceneTagCondition `json:"cond"`
}

type SceneTagCondition struct {
	CondType CondType `json:"condType"`
	Param1   int      `json:"param1"`
	Param2   int      `json:"param2"`
}

type CondType string

const (
	SCENE_TAG_COND_TYPE_ACTIVITY_CONTENT_OPEN  CondType = "SCENE_TAG_COND_TYPE_ACTIVITY_CONTENT_OPEN"
	SCENE_TAG_COND_TYPE_QUEST_FINISH           CondType = "SCENE_TAG_COND_TYPE_QUEST_FINISH"
	SCENE_TAG_COND_TYPE_QUEST_GLOBAL_VAR_EQUAL CondType = "SCENE_TAG_COND_TYPE_QUEST_GLOBAL_VAR_EQUAL"
	SCENE_TAG_COND_TYPE_SPECIFIC_ACTIVITY_OPEN CondType = "SCENE_TAG_COND_TYPE_SPECIFIC_ACTIVITY_OPEN"
)

func LoadSceneTagData() map[int]*TagData {
	data := []*TagData{}
	prefix := config.Conf.FolderStructure.Resources
	fp := filepath.Join(prefix + "ExcelBinOutput/" + SceneTagDataResource)
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

	sceneTagDataMap := make(map[int]*TagData)

	for _, v := range data {
		sceneTagDataMap[v.Id] = v
	}
	return sceneTagDataMap
}
