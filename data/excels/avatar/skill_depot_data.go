package avatar

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
)

const SkillDepotDataResource = "AvatarSkillDepotExcelConfigData.json"

type SkillDepotData struct {
	Id                      int                        `json:"id,omitempty"`
	EnergySkill             int                        `json:"energySkill,omitempty"`
	AttackModeSkill         int                        `json:"attackModeSkill,omitempty"`
	Skills                  []int                      `json:"skills,omitempty"`
	SubSkills               []int                      `json:"subSkills,omitempty"`
	ExtraAbilities          []string                   `json:"extraAbilities,omitempty"`
	Talents                 []int                      `json:"talents,omitempty"`
	InherentProudSkillOpens []*InherentProudSkillOpens `json:"inherentProudSkillOpens,omitempty"`
	TalentStarName          string                     `json:"talentStarName,omitempty"`
	SkillDepotAbilityGroup  string                     `json:"skillDepotAbilityGroup,omitempty"`

	// Transient
	EnergySkillData  *SkillData `json:"energySkillData,omitempty"`
	ElementType      any        `json:"elementType,omitempty"` // todo INCOMPLETE: ElementType
	Abilities        []int      `json:"abilities,omitempty"`
	TalentCostItemId int        `json:"talentCostItemId,omitempty"`
}

type InherentProudSkillOpens struct {
	ProudSkillGroupId      int `json:"proudSkillGroupId,omitempty"`
	NeedAvatarPromoteLevel int `json:"needAvatarPromoteLevel,omitempty"`
}

func LoadAvatarSkillDepotData() map[int]*SkillDepotData {
	data := []*SkillDepotData{}
	prefix := config.Conf.FolderStructure.Resources
	fp := filepath.Join(prefix + "ExcelBinOutput/" + SkillDepotDataResource)
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

	sceneTagDataMap := make(map[int]*SkillDepotData)

	for _, v := range data {
		sceneTagDataMap[v.Id] = v
	}
	return sceneTagDataMap
}
