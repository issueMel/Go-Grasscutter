package inventory

import (
	"Go-Grasscutter/generated/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GameItem struct {
	Id      primitive.ObjectID `bson:"id"`
	OwnerId int                `bson:"ownerId"`
	ItemId  int                `bson:"itemId"`
	Count   int                `bson:"count"`
	Guid    int64
	// itemdata

	// Equips
	Level        int  `bson:"level"`
	Exp          int  `bson:"exp"`
	TotalExp     int  `bson:"totalExp"`
	PromoteLevel int  `bson:"promoteLevel"`
	Locked       bool `bson:"locked"`

	// Weapon
	Affixes    []int `bson:"affixes"`
	Refinement int   `bson:"refinement"`

	// Relic
	MainPropId       int   `bson:"mainPropId"`
	AppendPropIdList []int `bson:"appendPropIdList"`

	EquipCharacter int `bson:"equipCharacter"`
	// weaponEntity
	NewItem bool
}

func (item *GameItem) ToProto() *pb.Item {
	prot := pb.Item{
		Guid:   uint64(item.Guid),
		ItemId: uint32(item.ItemId),
		// todo Detail
	}
	return &prot
}
