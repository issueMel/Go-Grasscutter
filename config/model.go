package config

type Config struct {
	FolderStructure *FolderStructure `json:"folderStructure,omitempty"`
	DatabaseInfo    *DatabaseInfo    `json:"databaseInfo,omitempty"`
	Language        *Language        `json:"language,omitempty"`
	Account         *Account         `json:"account,omitempty"`
	Server          *Server          `json:"server,omitempty"`
	DebugMode       *DebugMode       `json:"debugMode,omitempty"`
	Version         int              `json:"version,omitempty"`
}

type Server struct {
	DebugWhitelist []interface{} `json:"debugWhitelist,omitempty"`
	DebugBlacklist []interface{} `json:"debugBlacklist,omitempty"`
	RunMode        string        `json:"runMode,omitempty"`
	LogCommands    bool          `json:"logCommands,omitempty"`
	FastRequire    bool          `json:"fastRequire,omitempty"`
	Http           *Http         `json:"http,omitempty"`
	Game           *Game         `json:"game,omitempty"`
	Dispatch       *Dispatch     `json:"dispatch,omitempty"`
	DebugMode      *DebugMode    `json:"debugMode,omitempty"`
}

type FolderStructure struct {
	Resources string `json:"resources,omitempty"`
	Data      string `json:"data,omitempty"`
	Packets   string `json:"packets,omitempty"`
	Scripts   string `json:"scripts,omitempty"`
	Plugins   string `json:"plugins,omitempty"`
	Cache     string `json:"cache,omitempty"`
}

type DataStore struct {
	ConnectionUri string `json:"connectionUri"`
	Collection    string `json:"collection"`
}

type DatabaseInfo struct {
	Server *DataStore `json:"server,omitempty"`
	Game   *DataStore `json:"game,omitempty"`
}

type Language struct {
	Language string `json:"language,omitempty"`
	Fallback string `json:"fallback,omitempty"`
	Document string `json:"document,omitempty"`
}

type Account struct {
	AutoCreate          bool     `json:"autoCreate,omitempty"`
	ExperimentalRealPwd bool     `json:"EXPERIMENTAL_RealPassword,omitempty"`
	DefaultPermissions  []string `json:"defaultPermissions,omitempty"`
	MaxPlayer           int      `json:"maxPlayer,omitempty"`
}

type Http struct {
	StartImmediately bool        `json:"startImmediately,omitempty"`
	BindAddress      string      `json:"bindAddress,omitempty"`
	BindPort         int         `json:"bindPort,omitempty"`
	AccessAddress    string      `json:"accessAddress,omitempty"`
	AccessPort       int         `json:"accessPort,omitempty"`
	Encryption       *Encryption `json:"encryption,omitempty"`
	Policies         *Policies   `json:"policies,omitempty"`
	Files            *Files      `json:"files,omitempty"`
}

type Encryption struct {
	UseEncryption    bool   `json:"useEncryption,omitempty"`
	UseInRouting     bool   `json:"useInRouting,omitempty"`
	Keystore         string `json:"keystore,omitempty"`
	KeystorePassword string `json:"keystorePassword,omitempty"`
}

type Policies struct {
	Cors struct {
		Enabled        bool     `json:"enabled,omitempty"`
		AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	} `json:"cors"`
}

type Files struct {
	IndexFile string `json:"indexFile,omitempty"`
	ErrorFile string `json:"errorFile,omitempty"`
}

type Game struct {
	BindAddress                string           `json:"bindAddress,omitempty"`
	BindPort                   int              `json:"bindPort,omitempty"`
	AccessAddress              string           `json:"accessAddress,omitempty"`
	AccessPort                 int              `json:"accessPort,omitempty"`
	UseUniquePacketKey         bool             `json:"useUniquePacketKey,omitempty"`
	LoadEntitiesForPlayerRange int              `json:"loadEntitiesForPlayerRange,omitempty"`
	EnableScriptInBigWorld     bool             `json:"enableScriptInBigWorld,omitempty"`
	EnableConsole              bool             `json:"enableConsole,omitempty"`
	KcpInterval                int              `json:"kcpInterval,omitempty"`
	LogPackets                 string           `json:"logPackets,omitempty"`
	IsShowPacketPayload        bool             `json:"isShowPacketPayload,omitempty"`
	IsShowLoopPackets          bool             `json:"isShowLoopPackets,omitempty"`
	CacheSceneEntitiesEveryRun bool             `json:"cacheSceneEntitiesEveryRun,omitempty"`
	GameOptions                *GameOptions     `json:"gameOptions,omitempty"`
	JoinOptions                *JoinOptions     `json:"joinOptions,omitempty"`
	ServerAccount              *ServerAccount   `json:"serverAccount,omitempty"`
	VisionOptions              []*VisionOptions `json:"visionOptions,omitempty"`
}

type GameOptions struct {
	InventoryLimits  *InventoryLimits `json:"inventoryLimits,omitempty"`
	AvatarLimits     *AvatarLimits    `json:"avatarLimits,omitempty"`
	SceneEntityLimit int              `json:"sceneEntityLimit,omitempty"`
	WatchGachaConfig bool             `json:"watchGachaConfig,omitempty"`
	EnableShopItems  bool             `json:"enableShopItems,omitempty"`
	StaminaUsage     bool             `json:"staminaUsage,omitempty"`
	EnergyUsage      bool             `json:"energyUsage,omitempty"`
	FishhookTeleport bool             `json:"fishhookTeleport,omitempty"`
	TrialCostumes    bool             `json:"trialCostumes,omitempty"`
	Questing         *Questing        `json:"questing,omitempty"`
	ResinOptions     *ResinOptions    `json:"resinOptions,omitempty"`
	Rates            *Rates           `json:"rates,omitempty"`
	Handbook         *Handbook        `json:"handbook,omitempty"`
}

type JoinOptions struct {
	WelcomeEmotes  []int  `json:"welcomeEmotes,omitempty"`
	WelcomeMessage string `json:"welcomeMessage,omitempty"`
	WelcomeMail    struct {
		Title   string `json:"title,omitempty"`
		Content string `json:"content,omitempty"`
		Sender  string `json:"sender,omitempty"`
		Items   []struct {
			ItemId    int `json:"itemId,omitempty"`
			ItemCount int `json:"itemCount,omitempty"`
			ItemLevel int `json:"itemLevel,omitempty"`
		} `json:"items,omitempty"`
	} `json:"welcomeMail,omitempty"`
}

type ServerAccount struct {
	AvatarId      int    `json:"avatarId,omitempty"`
	NameCardId    int    `json:"nameCardId,omitempty"`
	AdventureRank int    `json:"adventureRank,omitempty"`
	WorldLevel    int    `json:"worldLevel,omitempty"`
	NickName      string `json:"nickName,omitempty"`
	Signature     string `json:"signature,omitempty"`
}

type VisionOptions []struct {
	Name        string `json:"name,omitempty"`
	VisionRange int    `json:"visionRange,omitempty"`
	GridWidth   int    `json:"gridWidth,omitempty"`
}

type AvatarLimits struct {
	SinglePlayerTeam int `json:"singlePlayerTeam,omitempty"`
	MultiplayerTeam  int `json:"multiplayerTeam,omitempty"`
}

type InventoryLimits struct {
	Weapons   int `json:"weapons"`
	Relics    int `json:"relics"`
	Materials int `json:"materials"`
	Furniture int `json:"furniture"`
	All       int `json:"all"`
}

type Questing struct {
	Enabled bool `json:"enabled,omitempty"`
}

type ResinOptions struct {
	ResinUsage   bool `json:"resinUsage"`
	Cap          int  `json:"cap"`
	RechargeTime int  `json:"rechargeTime"`
}

type Rates struct {
	AdventureExp float64 `json:"adventureExp,omitempty"`
	Mora         float64 `json:"mora,omitempty"`
	LeyLines     float64 `json:"leyLines,omitempty"`
}

type Handbook struct {
	Enable        bool `json:"enable,omitempty"`
	AllowCommands bool `json:"allowCommands,omitempty"`
	Limits        struct {
		Enabled     bool `json:"enabled,omitempty"`
		Interval    int  `json:"interval,omitempty"`
		MaxRequests int  `json:"maxRequests,omitempty"`
		MaxEntities int  `json:"maxEntities,omitempty"`
	} `json:"limits,omitempty"`
	Server struct {
		Enforced  bool   `json:"enforced,omitempty"`
		Address   string `json:"address,omitempty"`
		Port      int    `json:"port,omitempty"`
		CanChange bool   `json:"canChange,omitempty"`
	} `json:"server,omitempty"`
}

type Dispatch struct {
	Regions       []*Region `json:"regions"`
	DispatchUrl   string    `json:"dispatchUrl,omitempty"`
	EncryptionKey string    `json:"encryptionKey,omitempty"`
	DispatchKey   string    `json:"dispatchKey,omitempty"`
	DefaultName   string    `json:"defaultName"`
	LogRequests   string    `json:"logRequests,omitempty"`
}

type Region struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Ip    string `json:"ip"`
	Port  int    `json:"port"`
}

type DebugMode struct {
	ServerLoggerLevel struct {
		LevelInt int    `json:"levelInt"`
		LevelStr string `json:"levelStr"`
	} `json:"serverLoggerLevel"`
	ServicesLoggersLevel struct {
		LevelInt int    `json:"levelInt"`
		LevelStr string `json:"levelStr"`
	} `json:"servicesLoggersLevel"`
	LogPackets          string `json:"logPackets"`
	IsShowPacketPayload bool   `json:"isShowPacketPayload"`
	IsShowLoopPackets   bool   `json:"isShowLoopPackets"`
	LogRequests         string `json:"logRequests"`
}
