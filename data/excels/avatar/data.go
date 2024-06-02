package avatar

var avatarResources = []string{
	"AvatarExcelConfigData.json",
}

type AvatarData struct {
	IconName         string `json:"iconName,omitempty"`
	BodyType         string `json:"bodyType,omitempty"`
	QualityType      string `json:"qualityType,omitempty"`
	ChargeEfficiency int    `json:"chargeEfficiency,omitempty"`
	InitialWeapon    int    `json:"initialWeapon,omitempty"`
	// weaponType
	ImageName                    string `json:"imageName,omitempty"`
	AvatarPromoteId              int    `json:"avatarPromoteId,omitempty"`
	CutsceneShow                 string `json:"cutsceneShow,omitempty"`
	SkillDepotId                 int    `json:"skillDepotId,omitempty"`
	StaminaRecoverSpeed          int    `json:"staminaRecoverSpeed,omitempty"`
	CandSkillDepotIds            []int  `json:"candSkillDepotIds,omitempty"`
	AvatarIdentityType           string `json:"avatarIdentityType,omitempty"`
	AvatarPromoteRewardLevelList []int  `json:"avatarPromoteRewardLevelList,omitempty"`
	AvatarPromoteRewardIdList    []int  `json:"avatarPromoteRewardIdList,omitempty"`

	NameTextMapHash uint64  `json:"nameTextMapHash,omitempty"`
	HpBase          float32 `json:"hpBase,omitempty"`
	AttackBase      float32 `json:"attackBase,omitempty"`
	DefenseBase     float32 `json:"defenseBase,omitempty"`
	Critical        float32 `json:"critical,omitempty"`
	CriticalHurt    float32 `json:"criticalHurt,omitempty"`
	// propGrowCurves
}
