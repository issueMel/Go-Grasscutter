package avatar

type SkillDepotData struct {
	Id                      int
	EnergySkill             int
	AttackModeSkill         int
	Skills                  []int
	SubSkills               []int
	ExtraAbilities          []string
	Talents                 []int
	InherentProudSkillOpens []*InherentProudSkillOpens
	TalentStarName          string
	SkillDepotAbilityGroup  string

	// Transient
	EnergySkillData  *SkillData
	ElementType      any // todo INCOMPLETE: ElementType
	Abilities        []int
	TalentCostItemId int
}

type InherentProudSkillOpens struct {
	ProudSkillGroupId      int
	NeedAvatarPromoteLevel int
}
