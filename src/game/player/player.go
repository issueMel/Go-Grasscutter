package player

import (
	"Go-Grasscutter/src/game"
	game2 "Go-Grasscutter/src/server/game"
)

type Player struct {
	ID         int    `bson:"_id"`
	AccountID  string `bson:"accountId" index:"unique"`
	Account    *game.Account
	Session    *game2.GameSession
	SessionKey string `bson:"-"`
	Nickname   string `bson:"nickname"`
	Signature  string `bson:"signature"`
	HeadImage  int    `bson:"headImage"`
	NameCardID int    `bson:"nameCardId"`
	//Position         Position       `bson:"position"`
	//PrevPos          *Position      `bson:"prevPos"`
	//PrevPosForHome   *Position      `bson:"prevPosForHome"`
	PrevScene int `bson:"prevScene"`
	//Rotation         Position       `bson:"rotation"`
	Birthday PlayerBirthday `bson:"birthday"`
	//Codex            PlayerCodex    `bson:"codex"`
	ShowAvatars      bool             `bson:"showAvatars"`
	ShowAvatarList   []int            `bson:"showAvatarList"`
	ShowNameCardList []int            `bson:"showNameCardList"`
	Properties       map[int]int      `bson:"properties"`
	CurrentRealmID   int              `bson:"currentRealmId"`
	IsInEditMode     bool             `bson:"-"`
	WidgetID         int              `bson:"widgetId"`
	SceneID          int              `bson:"sceneId"`
	RegionID         int              `bson:"regionId"`
	MainCharacterID  int              `bson:"mainCharacterId"`
	InGodMode        bool             `bson:"inGodMode"`
	UnlimitedStamina bool             `bson:"unlimitedStamina"`
	NameCardList     map[int]struct{} `bson:"nameCardList"`
	CostumeList      map[int]struct{} `bson:"costumeList"`
}

func NewPlayer() *Player {
	return &Player{}
}
