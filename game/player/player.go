package player

import (
	"Go-Grasscutter/db"
	"Go-Grasscutter/game/achievement"
	"Go-Grasscutter/game/avatar"
	"Go-Grasscutter/game/battlepass"
	"Go-Grasscutter/game/city"
	"Go-Grasscutter/game/expedition"
	"Go-Grasscutter/game/friends"
	"Go-Grasscutter/game/gacha"
	"Go-Grasscutter/game/inventory"
	"Go-Grasscutter/game/managers/cooking"
	"Go-Grasscutter/game/managers/forging"
	"Go-Grasscutter/game/managers/mapmark"
	"Go-Grasscutter/game/quest"
	"Go-Grasscutter/game/shop"
	"Go-Grasscutter/game/world"
	"Go-Grasscutter/lib/kcp-go"
	"Go-Grasscutter/log"
	"Go-Grasscutter/utils"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

const collName = "players"

type Player struct {
	ID               int    `bson:"_id"`
	AccountID        string `bson:"accountId" index:"unique"`
	Account          *Account
	Session          *kcp.UDPSession
	SessionKey       string
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
	Codex            *PlayerCodex    `bson:"codex"`
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
	FlyCloakList              []int                                   `bson:"flyCloakList"`
	CostumeList               []int                                   `bson:"costumeList"`
	RewardedLevels            []int                                   `bson:"rewardedLevels"`
	HomeRewardedLevels        []int                                   `bson:"homeRewardedLevels"`
	RealmList                 []int                                   `bson:"realmList"`
	seenRealmList             []int                                   `bson:"seenRealmList"`
	UnlockedForgingBlueprints []int                                   `bson:"unlockedForgingBlueprints"`
	UnlockedCombines          []int                                   `bson:"unlockedCombines"`
	UnlockedFurniture         []int                                   `bson:"unlockedFurniture"`
	UnlockedFurnitureSuite    []int                                   `bson:"unlockedFurnitureSuite"`
	ExpeditionInfo            map[int64]*expedition.ExpeditionInfo    `bson:"expeditionInfo"`
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
	World    *world.World
	// curHomeWorld
	HasSentInitPacketInHome bool
	// scene
	WeatherId int
	// climate
	AreaId   int
	AreaType int

	// Player managers go here
	Avatars           *avatar.Storage
	Inventory         *inventory.Inventory
	QuestManager      *quest.Manager
	BattlePassManager *battlepass.Manager

	// Manager data (Save-able to the database)
	Achievements          *achievement.Achievements
	PlayerProfile         *friends.PlayerProfile   `bson:"playerProfile"`
	TeamManager           *TeamManager             `bson:"teamManager"`
	GachaInfo             *gacha.PlayerGachaInfo   `bson:"gachaInfo"`
	CollectionRecordStore *PlayerCollectionRecords `bson:"collectionRecordStore"`
	ShopLimit             []*shop.ShopLimit        `bson:"shopLimit"`

	// todo INCOMPLETE: home

	MoonCard bool `bson:"moonCard"`
	// moonCardStartTime time.Duration
	MoonCardDuration int `bson:"moonCardDuration"`
	// moonCardGetTimes []int `bson:"moonCardGetTimes"`

	Paused bool
	// QueuedTeleport
	EnterSceneToken       int
	SceneLoadState        SceneLoadState
	HasSentLoginPackets   bool
	NextSendPlayerLocTime int64
	// EnterHomeRequests map[int]any

	SpringLastUsed       int64                       `bson:"springLastUsed"`
	MapMarks             map[string]*mapmark.MapMark `bson:"mapMarks"`
	NextResinRefresh     int                         `bson:"nextResinRefresh"`
	ResinBuyCount        int                         `bson:"resinBuyCount"`
	LastDailyReset       int                         `bson:"lastDailyReset"`
	PlayerGameTime       int64                       `bson:"playerGameTime"`
	PlayerProgress       *PlayerProgress             `bson:"playerProgress"`
	ActiveQuestTimers    []int                       `bson:"activeQuestTimers"`
	MainCharacterElement string                      `bson:"mainCharacterElement"`
	CityInfoData         *city.CityInfoData          `bson:"cityInfoData"`
}

func NewPlayer(session *kcp.UDPSession) {

}

// Create
func CreatePlayer(account *Account, sess *kcp.UDPSession) *Player {
	p := &Player{
		Account:   account,
		AccountID: account.ID,
		Session:   sess,
		Nickname:  "Traveler",
		Signature: "",
	}
	// todo INCOMPLETE: applyProperties
	// todo INCOMPLETE: applyStartingSceneTags

	return p
}

func (p *Player) LoadFromDatabase() {
	// todo INCOMPLETE: init Manager
	// Make sure these exist
	if p.TeamManager == nil {
		p.TeamManager = &TeamManager{}
	}

	if p.Codex == nil {
		p.Codex = &PlayerCodex{}
	}

	if p.PlayerProfile == nil {
		p.PlayerProfile = &friends.PlayerProfile{}
	}

	if p.PlayerProfile == nil || p.PlayerProfile.Uid == 0 {
		// syncWithCharacter
		p.PlayerProfile = &friends.PlayerProfile{
			Uid:       p.ID,
			NameCard:  p.NameCardID,
			AvatarId:  p.HeadImage,
			Name:      p.Nickname,
			Signature: p.Signature,
			//PlayerLevel:      ,
			//WorldLevel:       ,
			LastActiveTime:  utils.GetCurrentSeconds(),
			EnterHomeOption: 0, // todo INCOMPLETE
		}
	}
	// todo INCOMPLETE: Load all Player managers From Database
	// Load from db
	wait := &sync.WaitGroup{}

	utils.DoAllFunc(wait,
		func() {
			p.Achievements = achievement.GetByPlayerUid(p.ID)
		},
		// p.Avatars.LoadFromDatabase,
	)

	// Wait for all tasks to finish.
	wait.Wait()
}

func (p *Player) OnLogin() {
	if p.SceneTags == nil || len(p.SceneTags) == 0 {
		// todo INCOMPLETE: applyStartingSceneTags()
	}

	// todo GameHome

	// todo Create world

	// todo Multiplayer setting

	// todo Execute daily reset logic if this is a new day.

	// todo Rewind active quests, and put the player to a rewind position it finds (if any) of an active quest

}

func GetPlayerByAccount(account *Account) *Player {
	p := &Player{}
	err := db.DB.Collection(collName).FindOne(context.Background(), bson.D{{"accountId", account.ID}}).Decode(p)
	if err != nil {
		if errors.Is(mongo.ErrNoDocuments, err) {
			return nil
		}
		log.SugaredLogger.Error(err)
		return nil
	}

	// init management
	p.ManagementInit()

	return p
}

// todo CHECK: move to NewPlayer()
func (p *Player) ManagementInit() {
	p.Avatars = &avatar.Storage{
		Uid:         p.ID,
		Avatars:     make(map[int]*avatar.Avatar),
		AvatarsGuid: make(map[int64]*avatar.Avatar),
	}

	p.Inventory = &inventory.Inventory{
		Uid:            p.ID,
		Store:          make(map[int64]*inventory.GameItem),
		InventoryTypes: make(map[int]*inventory.InventoryTab),
	}

	p.QuestManager = &quest.Manager{
		MainQuests:          make(map[int]*quest.GameMainQuest),
		AcceptProgressLists: make(map[int][]int),
		LoggedQuests:        make([]int, 0),
	}

	p.BattlePassManager = &battlepass.Manager{
		OwnerUid:     p.ID,
		Missions:     make(map[int]*battlepass.Mission),
		TakenRewards: make(map[int]*battlepass.Reward),
	}

	p.World = &world.World{} // todo INCOMPLETE
}

func CheckIfExists(uid int) bool {
	err := db.DB.Collection(collName).FindOne(context.Background(), bson.D{{"_id", uid}}).Err()
	if err == nil {
		return true
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		log.SugaredLogger.Error(err)
		return false
	}
	return false
}

func GetNextPlayerId(reservedId int) int {
	var id int
	if reservedId > 0 && CheckIfExists(reservedId) {
		id = reservedId
	} else {
		for CheckIfExists(id) {
			id = db.GetNextId("Player")
		}
	}
	return id
}

func (p *Player) GeneratePlayerUid(reservedId int) {
	var id int
	if reservedId > 0 && !CheckIfExists(reservedId) {
		id = reservedId
	} else {
		for CheckIfExists(id) {
			id = db.GetNextId("Player")
		}
	}
	p.ID = id
	go p.Save()
}

func (p *Player) Save() {
	_, err := db.DB.Collection(collName).InsertOne(context.Background(), p)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
}
