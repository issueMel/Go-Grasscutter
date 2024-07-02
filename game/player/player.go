package player

import (
	"Go-Grasscutter/data"
	"Go-Grasscutter/db"
	"Go-Grasscutter/game/achievement"
	"Go-Grasscutter/game/activity"
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
	"Go-Grasscutter/game/pros"
	"Go-Grasscutter/game/quest"
	"Go-Grasscutter/game/shop"
	"Go-Grasscutter/game/tunnel"
	"Go-Grasscutter/game/world"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/utils"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

const collName = "players"

type Player struct {
	ID               int    `bson:"_id"`
	AccountID        string `bson:"accountId"`
	Account          *Account
	Tunnel           *tunnel.Tunnel
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
	ShowNameCardList []int32         `bson:"showNameCardList"`
	Properties       map[int]int     `bson:"properties"` // int, pros.PlayerProperty
	CurrentRealmID   int             `bson:"currentRealmId"`
	IsInEditMode     bool
	WidgetID         int  `bson:"widgetId"`
	SceneID          int  `bson:"sceneId"`
	RegionID         int  `bson:"regionId"`
	MainCharacterID  int  `bson:"mainCharacterId"`
	InGodMode        bool `bson:"inGodMode"`
	UnlimitedStamina bool `bson:"unlimitedStamina"`

	NameCardList              []int                                   `bson:"nameCardList"`
	FlyCloakList              []int32                                 `bson:"flyCloakList"`
	CostumeList               []int32                                 `bson:"costumeList"`
	RewardedLevels            []int32                                 `bson:"rewardedLevels"`
	HomeRewardedLevels        []int                                   `bson:"homeRewardedLevels"`
	RealmList                 []int                                   `bson:"realmList"`
	seenRealmList             []int                                   `bson:"seenRealmList"`
	UnlockedForgingBlueprints []int                                   `bson:"unlockedForgingBlueprints"`
	UnlockedCombines          []int32                                 `bson:"unlockedCombines"`
	UnlockedFurniture         []int32                                 `bson:"unlockedFurniture"`
	UnlockedFurnitureSuite    []int32                                 `bson:"unlockedFurnitureSuite"`
	ExpeditionInfo            map[int64]*expedition.ExpeditionInfo    `bson:"expeditionInfo"`
	UnlockedRecipies          map[int]int                             `bson:"unlockedRecipies"`
	ActiveForges              []*forging.ActiveForgeData              `bson:"activeForges"`
	ActiveCookCompounds       map[int]*cooking.ActiveCookCompoundData `bson:"activeCookCompounds"`
	QuestGlobalVariables      map[int]int                             `bson:"questGlobalVariables"`
	OpenStates                map[int]int                             `bson:"openStates"`
	SceneTags                 map[int][]int                           `bson:"sceneTags"`
	UnlockedSceneAreas        map[int][]int32                         `bson:"unlockedSceneAreas"`
	UnlockedScenePoints       map[int][]int32                         `bson:"unlockedScenePoints"`
	ChatEmojiIdList           []int32                                 `bson:"chatEmojiIdList"`

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
	ActivityManager   *activity.Manager

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

	ClientAbility []*pb.AbilityInvokeEntry

	SpringLastUsed       int64                       `bson:"springLastUsed"`
	MapMarks             map[string]*mapmark.MapMark `bson:"mapMarks"`
	NextResinRefresh     int                         `bson:"nextResinRefresh"`
	ResinBuyCount        int                         `bson:"resinBuyCount"`
	LastDailyReset       int                         `bson:"lastDailyReset"`
	PlayerGameTime       int64                       `bson:"playerGameTime"`
	PlayerProgress       *PlayerProgress             `bson:"playerProgress"`
	ActiveQuestTimers    []int                       `bson:"activeQuestTimers"`
	MainCharacterElement string                      `bson:"mainCharacterElement"`
	CityInfoData         map[int]*city.CityInfoData  `bson:"cityInfoData"`
}

func NewPlayer() *Player {
	return &Player{}
}

// Create
func CreatePlayer(account *Account, tunnel *tunnel.Tunnel) *Player {
	p := &Player{
		Account:   account,
		AccountID: account.ID,
		Tunnel:    tunnel,
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

	if p.PlayerProfile == nil || p.PlayerProfile.Uid == 0 {
		// syncWithCharacter
		p.PlayerProfile = &friends.PlayerProfile{
			Uid:             p.ID,
			NameCard:        p.NameCardID,
			AvatarId:        p.HeadImage,
			Name:            p.Nickname,
			Signature:       p.Signature,
			PlayerLevel:     p.GetLevel(),
			WorldLevel:      p.GetWorldLevel(),
			LastActiveTime:  utils.GetCurrentSeconds(),
			EnterHomeOption: 1, // todo INCOMPLETE
		}

	}
	// todo INCOMPLETE: Load all Player managers From Database
	// Load from db
	wait := &sync.WaitGroup{}

	utils.DoAllFunc(wait,
		func() {
			p.Achievements = achievement.GetByPlayerUid(p.ID)
		},
		p.Avatars.LoadFromDatabase,
		p.Inventory.LoadFromDatabase,
		// getFriendsList
		// getMailHandler
		p.QuestManager.LoadFromDatabase,
		p.BattlePassManager.LoadFromDatabase,
	)

	// Wait for all tasks to finish.
	wait.Wait()
}

func (p *Player) OnLogin() {
	// Ensure the player has valid scene tags, allows old accounts to work
	if p.SceneTags == nil || len(p.SceneTags) == 0 {
		// p.applyStartingSceneTags()
	}

	// todo GameHome

	// todo Create world

	// todo Multiplayer setting

	// todo Execute daily reset logic if this is a new day.

	// todo Rewind active quests, and put the player to a rewind position it finds (if any) of an active quest

}

// TODO ENHANCE: improve applyStartingSceneTags performance
// Use DB addToSet ?
func (p *Player) applyStartingSceneTags() {
	for _, sceneTag := range data.GameData.SceneTagDataMap {
		if sceneTag.IsDefaultValid {
			set := make(map[int]struct{})
			if _, ok := p.SceneTags[sceneTag.SceneId]; ok {
				for _, v := range p.SceneTags[sceneTag.SceneId] {
					set[v] = struct{}{}
				}
			}

			set[sceneTag.Id] = struct{}{}

			ints := make([]int, len(set))
			for i := range set {
				ints = append(ints, i)
			}

			p.SceneTags[sceneTag.SceneId] = ints
		}
	}
}

func GetPlayerByAccount(account *Account) *Player {
	p := NewPlayer()
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
		Uid:                 p.ID,
		MainQuests:          make(map[int]*quest.GameMainQuest),
		AcceptProgressLists: make(map[int][]int),
		LoggedQuests:        make([]int, 0),
	}

	p.BattlePassManager = &battlepass.Manager{
		OwnerUid:     p.ID,
		Missions:     make(map[int]*battlepass.Mission),
		TakenRewards: make(map[int]*battlepass.Reward),
	}

	p.InitWorld()
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
	opts := options.Update().SetUpsert(true)
	_, err := db.DB.Collection(collName).UpdateOne(context.Background(), bson.D{{"id", p.ID}}, p, opts)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}
}

func (p *Player) GetSocialDetail() *pb.SocialDetail {
	infoList := make([]*pb.SocialShowAvatarInfo, 0)
	if p.isOnline() {
		if len(p.ShowAvatarList) > 0 {
			infoList = make([]*pb.SocialShowAvatarInfo, 0, len(p.ShowAvatarList))
			for _, avatarId := range p.ShowAvatarList {
				avatar, ok := p.Avatars.Avatars[avatarId]
				if ok {
					infoList = append(infoList, &pb.SocialShowAvatarInfo{
						AvatarId:  uint32(avatarId),
						Level:     uint32(avatar.Level),
						CostumeId: uint32(avatar.Costume),
					})
				}

			}
		}
	} else {
		// todo !player.isOnline()
	}
	prot := &pb.SocialDetail{
		Uid: uint32(p.ID),
		ProfilePicture: &pb.ProfilePicture{
			AvatarId: uint32(p.HeadImage),
		},
		Nickname:              p.Nickname,
		Signature:             p.Signature,
		Level:                 uint32(p.GetLevel()),
		Birthday:              p.Birthday.GetFilledProtoWhenNotEmpty(),
		WorldLevel:            uint32(p.GetWorldLevel()),
		NameCardId:            uint32(p.NameCardID),
		IsShowAvatar:          p.ShowAvatars,
		ShowAvatarInfoList:    infoList,
		ShowNameCardIdList:    utils.ToUint32Slice(p.ShowNameCardList),
		FinishAchievementNum:  uint32(p.Achievements.FinishedAchievementNum),
		FriendEnterHomeOption: 0, // todo game home
	}

	return prot
}

func (p *Player) GetProperty(id int) int {
	if p.Properties != nil {
		return p.Properties[id]
	}
	return 0
}

func (p *Player) GetLevel() int {
	return p.GetProperty(pros.PROP_PLAYER_LEVEL)
}

func (p *Player) isOnline() bool {
	// todo session.isActive()
	return p.Tunnel != nil
}
