package player

import (
	"Go-Grasscutter/db"
	"Go-Grasscutter/game/city"
	"Go-Grasscutter/game/expedition"
	"Go-Grasscutter/game/friends"
	"Go-Grasscutter/game/gacha"
	"Go-Grasscutter/game/managers/cooking"
	"Go-Grasscutter/game/managers/forging"
	"Go-Grasscutter/game/managers/mapmark"
	"Go-Grasscutter/game/shop"
	"Go-Grasscutter/game/world"
	"Go-Grasscutter/lib/kcp-go"
	"Go-Grasscutter/log"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Player struct {
	ID               int    `bson:"_id"`
	AccountID        string `bson:"accountId" index:"unique"`
	Account          *Account
	Session          *kcp.UDPSession
	SessionKey       string          `bson:"-"`
	Nickname         string          `bson:"nickname"`
	Signature        string          `bson:"signature"`
	HeadImage        int             `bson:"headImage"`
	NameCardID       int             `bson:"nameCardId"`
	Position         *world.Position `bson:"position"`
	PrevPos          *world.Position `bson:"prevPos"`
	PrevPosForHome   *world.Position `bson:"prevPosForHome"`
	PrevScene        int             `bson:"prevScene"`
	Rotation         *world.Position `bson:"rotation"`
	Birthday         *PlayerBirthday `bson:"birthday"`
	Codex            PlayerCodex     `bson:"codex"`
	ShowAvatars      bool            `bson:"showAvatars"`
	ShowAvatarList   []int           `bson:"showAvatarList"`
	ShowNameCardList []int           `bson:"showNameCardList"`
	Properties       map[int]int     `bson:"properties"`
	CurrentRealmID   int             `bson:"currentRealmId"`
	IsInEditMode     bool
	WidgetID         int  `bson:"widgetId"`
	SceneID          int  `bson:"sceneId"`
	RegionID         int  `bson:"regionId"`
	MainCharacterID  int  `bson:"mainCharacterId"`
	InGodMode        bool `bson:"inGodMode"`
	UnlimitedStamina bool `bson:"unlimitedStamina"`

	NameCardList              []int                                   `bson:"nameCardList"`
	CostumeList               []int                                   `bson:"costumeList"`
	RewardedLevels            []int                                   `bson:"rewardedLevels"`
	HomeRewardedLevels        []int                                   `bson:"homeRewardedLevels"`
	RealmList                 []int                                   `bson:"realmList"`
	seenRealmList             []int                                   `bson:"seenRealmList"`
	UnlockedForgingBlueprints []int                                   `bson:"unlockedForgingBlueprints"`
	UnlockedCombines          []int                                   `bson:"unlockedCombines"`
	UnlockedFurniture         []int                                   `bson:"unlockedFurniture"`
	UnlockedFurnitureSuite    []int                                   `bson:"unlockedFurnitureSuite"`
	ExpeditionInfo            map[int64]*expedition.ExpeditionInfo    `bson:"expeditionInfo"` // todo int64, ExpeditionInfo
	UnlockedRecipies          map[int]int                             `bson:"unlockedRecipies"`
	ActiveForges              []*forging.ActiveForgeData              `bson:"activeForges"`
	ActiveCookCompounds       map[int]*cooking.ActiveCookCompoundData `bson:"activeCookCompounds"`
	QuestGlobalVariables      map[int]int                             `bson:"questGlobalVariables"`
	OpenStates                map[int]int                             `bson:"openStates"`
	SceneTags                 map[int][]int                           `bson:"sceneTags"`
	UnlockedSceneAreas        map[int][]int                           `bson:"unlockedSceneAreas"`
	UnlockedScenePoints       map[int][]int                           `bson:"unlockedScenePoints"`
	ChatEmojiIdList           []int                                   `bson:"chatEmojiIdList"`

	NextGuid int64
	PeerId   int

	// Player managers go here

	// Manager data (Save-able to the database)
	// achievements todo
	PlayerProfile         *friends.PlayerProfile   `bson:"playerProfile"`
	TeamManager           *TeamManager             `bson:"teamManager"`
	GachaInfo             *gacha.PlayerGachaInfo   `bson:"gachaInfo"`
	CollectionRecordStore *PlayerCollectionRecords `bson:"collectionRecordStore"`
	ShopLimit             []*shop.ShopLimit        `bson:"shopLimit"`

	MoonCard          bool                        `bson:"moonCard"`
	MoonCardDuration  int                         `bson:"moonCardDuration"`
	SpringLastUsed    int64                       `bson:"springLastUsed"`
	MapMarks          map[string]*mapmark.MapMark `bson:"mapMarks"`
	NextResinRefresh  int                         `bson:"nextResinRefresh"`
	ResinBuyCount     int                         `bson:"resinBuyCount"`
	LastDailyReset    int                         `bson:"lastDailyReset"`
	PlayerGameTime    int64                       `bson:"playerGameTime"`
	PlayerProgress    *PlayerProgress             `bson:"playerProgress"`
	ActiveQuestTimers []int                       `bson:"activeQuestTimers"`
	// todo mainCharacterElement
	cityInfoData *city.CityInfoData `bson:"cityInfoData"`
}

func NewPlayer() *Player {
	return &Player{}
}

func LoadPlayer(id int) *Player {
	p := &Player{}
	err := db.DB.Collection("players").FindOne(context.Background(), bson.D{{"_id", id}}).Decode(p)
	if err != nil {
		log.SugaredLogger.Error(err)
	}
	return p
}
