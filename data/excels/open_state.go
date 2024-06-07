package excels

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
)

const OpenStateConfigResource = "OpenStateConfigData.json"

type OpenStateData struct {
	Id              int  `json:"id,omitempty"`
	DefaultState    bool `json:"defaultState,omitempty"`
	AllowClientOpen bool `json:"allowClientOpen,omitempty"`
	SystemOpenUiId  int  `json:"systemOpenUiId,omitempty"`
}

type OpenStateCond struct {
	CondType string `json:"condType,omitempty"`
	Param    int    `json:"param,omitempty"`
	Param2   int    `json:"param2,omitempty"`
}

func LoadOpenState() []*OpenStateData {
	prefix := config.Conf.FolderStructure.Resources
	fp := filepath.Join(prefix + "ExcelBinOutput/" + OpenStateConfigResource)
	file, err := os.ReadFile(fp)
	if err != nil {
		log.SugaredLogger.Panic(err)
		return nil
	}
	list := []*OpenStateData{}

	err = json.Unmarshal(file, &list)
	if err != nil {
		log.SugaredLogger.Panic(err)
		return nil
	}

	return list
}
