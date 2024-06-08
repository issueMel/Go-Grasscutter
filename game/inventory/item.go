package inventory

import (
	"Go-Grasscutter/data/excels"
	"Go-Grasscutter/db"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collName = "items"

type GameItem struct {
	Id       primitive.ObjectID `bson:"id"`
	OwnerId  int                `bson:"ownerId"`
	ItemId   int                `bson:"itemId"`
	Count    int                `bson:"count"`
	Guid     int64
	ItemData *excels.ItemData

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
		// todo INCOMPLETE: Detail
	}
	return &prot
}

func GetInventoryItems(uid int) []*GameItem {
	find, err := db.DB.Collection(collName).Find(context.Background(), bson.D{{"ownerId", uid}})
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}
	var items []*GameItem
	err = find.All(context.Background(), &items)
	if err != nil {
		log.SugaredLogger.Error(err)
		return nil
	}
	return items
}
