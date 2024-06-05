package avatar

import (
	"Go-Grasscutter/generated/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Avatar struct {
	// Equips map[int]GameItem
	FightProperties    map[int]float32
	FightPropOverrides map[int]float32
	Id                 primitive.ObjectID `bson:"_id"`
	OwnerId            int                `bson:"ownerId"`
	// owner player.Player
	// avatarData AvatarData
	// skillDepot
	Guid                int64
	AvatarId            int     `bson:"avatarId"`
	Level               int     `bson:"level"`
	Exp                 int     `bson:"exp"`
	PromoteLevel        int     `bson:"promoteLevel"`
	Satiation           int     `bson:"satiation"`
	SatiationPenalty    int     `bson:"satiationPenalty"`
	CurrentHp           float32 `bson:"currentHp"`
	CurrentEnergy       float32 `bson:"currentEnergy"`
	ExtraAbilityEmbryos map[string]struct{}
	Fetters             []int       `bson:"fetters"`
	SkillLevelMap       map[int]int `bson:"skillLevelMap"`
	SkillExtraChargeMap map[int]int
	ProudSkillBonusMap  map[int]int
	SkillDepotId        int   `bson:"skillDepotId"`
	TalentIdList        []int `bson:"talentIdList"`   // set<int>
	ProudSkillList      []int `bson:"proudSkillList"` // set<int>
	FlyCloak            int   `bson:"flyCloak"`
	Costume             int   `bson:"costume"`
	BornTime            int   `bson:"bornTime"`
	FetterLevel         int   `bson:"fetterLevel"`
	FetterExp           int   `bson:"fetterExp"`
	NameCardRewardId    int   `bson:"nameCardRewardId"`
	NameCardId          int   `bson:"nameCardId"`
	TrialAvatarId       int   `bson:"trialAvatarId"` // trial avatar property
	GrantReason         int   `bson:"grantReason"`   // cannot store to db if grant reason is not integer
	FromParentQuestId   int   `bson:"fromParentQuestId"`
	// so far no outer class or prop value has information of this, but from packet:
	// 1 = normal, 2 = trial avatar
	AvatarType int
}

func (a *Avatar) ToProto() *pb.AvatarInfo {
	// todo INCOMPLETE: ToProto()
	return &pb.AvatarInfo{}
}
