package avatar

type SkillData struct {
	Id                 int
	cdTime             float32
	CostElemVal        int
	MaxChargeNum       int
	TriggerID          int
	IsAttackCameraLock bool
	ProudSkillGroupId  int
	CostElemType       any // todo INCOMPLETE: ElementType
	NameTextMapHash    int64
	DescTextMapHash    int64
	AbilityName        string
}
