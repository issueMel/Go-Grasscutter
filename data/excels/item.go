package excels

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/data/common"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
)

var itemResources = []string{
	"MaterialExcelConfigData.json",
	"WeaponExcelConfigData.json",
	"ReliquaryExcelConfigData.json",
	"HomeWorldFurnitureExcelConfigData.json",
}

type ItemData struct {
	Id                         int    `json:"id,omitempty"`
	StackLimit                 int    `json:"stackLimit,omitempty"`
	MaxUseCount                int    `json:"maxUseCount,omitempty"`
	RankLevel                  int    `json:"rankLevel,omitempty"`
	EffectName                 string `json:"effectName,omitempty"`
	Rank                       int    `json:"rank,omitempty"`
	Weight                     int    `json:"weight,omitempty"`
	GadgetId                   int    `json:"gadgetId,omitempty"`
	Icon                       string `json:"icon,omitempty"`
	DestroyReturnMaterial      []int  `json:"destroyReturnMaterial,omitempty"`
	DestroyReturnMaterialCount []int  `json:"destroyReturnMaterialCount,omitempty"`

	// Enums
	ItemType     string `json:"itemType,omitempty"`
	MaterialType string `json:"materialType,omitempty"`
	EquipType    string `json:"equipType,omitempty"`
	EffectType   string `json:"effectType,omitempty"`
	DestroyRule  string `json:"destroyRule,omitempty"`

	// Food
	FoodQuality     string `json:"foodQuality,omitempty"`
	SatiationParams []int  `json:"satiationParams,omitempty"`

	// Usable item
	UseTarget string                `json:"useTarget,omitempty"` // enum ItemUseTarget
	ItemUse   []*common.ItemUseData `json:"itemUse,omitempty"`   // todo ItemUseData
	// itemUseActions
	UseOnGain bool `json:"useOnGain,omitempty"`

	// Relic
	MainPropDepotId   int   `json:"mainPropDepotId,omitempty"`
	AppendPropDepotId int   `json:"appendPropDepotId,omitempty"`
	AppendPropNum     int   `json:"appendPropNum,omitempty"`
	SetId             int   `json:"setId,omitempty"`
	AddPropLevels     []int `json:"addPropLevels,omitempty"`
	BaseConvExp       int   `json:"baseConvExp,omitempty"`
	MaxLevel          int   `json:"maxLevel,omitempty"`

	// Weapon
	WeaponPromoteId int               `json:"weaponPromoteId,omitempty"`
	WeaponBaseExp   int               `json:"weaponBaseExp,omitempty"`
	StoryId         int               `json:"storyId,omitempty"`
	AvatarPromoteId int               `json:"avatarPromoteId,omitempty"`
	AwakenMaterial  int               `json:"awakenMaterial,omitempty"`
	AwakenCosts     []int             `json:"awakenCosts,omitempty"`
	SkillAffix      []int             `json:"skillAffix,omitempty"`
	WeaponProp      []*WeaponProperty `json:"weaponProp,omitempty"`

	// Hash
	NameTextMapHash int64 `json:"nameTextMapHash,omitempty"`

	// Furniture
	Comfort              int    `json:"comfort,omitempty"`
	FurnType             []int  `json:"furnType,omitempty"`
	FurnitureGadgetID    []int  `json:"furnitureGadgetID,omitempty"`
	SpecialFurnitureType string `json:"specialFurnitureType,omitempty"` // enum SpecialFurnitureType
	RoomSceneId          int    `json:"roomSceneId,omitempty"`

	// addPropLevelSet map[int]struct{}
}

type WeaponProperty struct {
	PropType  string  `json:"propType,omitempty"` // enum FightProperty
	InitValue float32 `json:"initValue,omitempty"`
	Type      string  `json:"type,omitempty"`
}

func LoadItemData() map[int]*ItemData {
	prefix := config.Conf.FolderStructure.Resources
	m := make(map[int]*ItemData)
	for _, fileName := range itemResources {
		fp := filepath.Join(prefix + "ExcelBinOutput/" + fileName)
		file, err := os.ReadFile(fp)
		if err != nil {
			log.SugaredLogger.Panic(err)
			continue
		}
		list := []*ItemData{}

		err = json.Unmarshal(file, &list)
		if err != nil {
			log.SugaredLogger.Panic(err)
			continue
		}

		// todo INCOMPLETE: onLoad()

		for _, item := range list {
			m[item.Id] = item
		}
	}
	return m
}
