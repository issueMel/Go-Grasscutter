package avatar

import (
	"Go-Grasscutter/data/excels/avatar"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Avatar struct {
	// Equips map[int]inventory.GameItem
	FightProperties    map[int]float32
	FightPropOverrides map[int]float32
	Id                 primitive.ObjectID `bson:"_id"`
	OwnerId            int                `bson:"ownerId"`
	// owner player.Player
	AvatarData          *avatar.Data
	SkillDepot          *avatar.SkillDepotData
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
	avatarFetter := pb.AvatarFetterInfo{
		ExpLevel:                uint32(a.FetterLevel),
		FetterList:              make([]*pb.FetterData, len(a.Fetters)),
		RewardedFetterLevelList: make([]uint32, 0),
	}
	if a.FetterLevel != 10 {
		avatarFetter.ExpNumber = uint32(a.FetterExp)
	}

	if len(a.Fetters) > 0 {
		for i := range a.Fetters {
			avatarFetter.FetterList[i] = &pb.FetterData{
				FetterId:    uint32(a.Fetters[i]),
				FetterState: 3, // todo enum FetterState
			}
		}
	}

	// todo NameCardList
	// todo RewardedFetterLevelList
	proto := &pb.AvatarInfo{
		AvatarId:                uint32(a.AvatarId),
		Guid:                    uint64(a.Guid),
		LifeState:               1,
		TalentIdList:            make([]uint32, 0),
		FightPropMap:            make(map[uint32]float32),
		SkillDepotId:            uint32(a.SkillDepotId),
		CoreProudSkillLevel:     6, // todo getCoreProudSkillLevel()
		SkillLevelMap:           make(map[uint32]uint32),
		InherentProudSkillList:  make([]uint32, 0),
		ProudSkillExtraLevelMap: make(map[uint32]uint32),
		AvatarType:              uint32(a.AvatarType),
		BornTime:                uint32(a.BornTime),
		FetterInfo:              &avatarFetter,
		WearingFlycloakId:       uint32(a.FlyCloak),
		CostumeId:               uint32(a.Costume),
	}

	err := copier.Copy(&proto.TalentIdList, a.TalentIdList)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	err = copier.Copy(&proto.FightPropMap, a.FightProperties)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	err = copier.Copy(&proto.SkillLevelMap, a.SkillLevelMap)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	err = copier.Copy(&proto.InherentProudSkillList, a.ProudSkillList)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}

	err = copier.Copy(&proto.ProudSkillExtraLevelMap, a.ProudSkillBonusMap)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}
	return proto
}
