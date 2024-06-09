package avatar

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
)

const avatarResources = "AvatarExcelConfigData.json"

type Data struct {
	IconName                     string  `json:"iconName,omitempty"`
	BodyType                     string  `json:"bodyType,omitempty"`
	QualityType                  string  `json:"qualityType,omitempty"`
	ChargeEfficiency             float32 `json:"chargeEfficiency,omitempty"`
	InitialWeapon                int     `json:"initialWeapon,omitempty"`
	WeaponType                   string  `json:"weaponType"` // todo WeaponType
	ImageName                    string  `json:"imageName,omitempty"`
	AvatarPromoteId              int     `json:"avatarPromoteId,omitempty"`
	CutsceneShow                 string  `json:"cutsceneShow,omitempty"`
	SkillDepotId                 int     `json:"skillDepotId,omitempty"`
	StaminaRecoverSpeed          float32 `json:"staminaRecoverSpeed,omitempty"`
	CandSkillDepotIds            []int   `json:"candSkillDepotIds,omitempty"`
	AvatarIdentityType           string  `json:"avatarIdentityType,omitempty"`
	AvatarPromoteRewardLevelList []int   `json:"avatarPromoteRewardLevelList,omitempty"`
	AvatarPromoteRewardIdList    []int   `json:"avatarPromoteRewardIdList,omitempty"`

	NameTextMapHash int64 `json:"nameTextMapHash,omitempty"`

	HpBase         float32          `json:"hpBase,omitempty"`
	AttackBase     float32          `json:"attackBase,omitempty"`
	DefenseBase    float32          `json:"defenseBase,omitempty"`
	Critical       float64          `json:"critical,omitempty"`
	CriticalHurt   float32          `json:"criticalHurt,omitempty"`
	PropGrowCurves []*PropGrowCurve `json:"propGrowCurves,omitempty"`
	Id             int              `json:"id,omitempty"`
}

type PropGrowCurve struct {
	Type      string `json:"type,omitempty"`
	GrowCurve string `json:"growCurve,omitempty"`
}

func LoadAvatarData() map[int]*Data {
	data := []*Data{}
	prefix := config.Conf.FolderStructure.Resources
	fp := filepath.Join(prefix + "ExcelBinOutput/" + avatarResources)
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
		sceneTagDataMap[v.Id] = v
	}
	return sceneTagDataMap
}
