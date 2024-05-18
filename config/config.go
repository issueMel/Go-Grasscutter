package config

type Config struct {
	FolderStructure *FolderStructure `json:"folderStructure"`
	DatabaseInfo    *DatabaseInfo    `json:"databaseInfo"`
	Language        *Language        `json:"language"`
	Account         *Account         `json:"account"`
	Server          *Server          `json:"server"`
	Dispatch        *Dispatch        `json:"dispatch"`
	DebugMode       *DebugMode       `json:"debugMode"`
	Version         int              `json:"version"`
}

type Server struct {
	DebugWhitelist []interface{} `json:"debugWhitelist"`
	DebugBlacklist []interface{} `json:"debugBlacklist"`
	RunMode        string        `json:"runMode"`
	LogCommands    bool          `json:"logCommands"`
	FastRequire    bool          `json:"fastRequire"`
	Http           *Http         `json:"http"`
	Game           *Game         `json:"game"`
	Dispatch       *Dispatch     `json:"dispatch"`
	DebugMode      *DebugMode    `json:"debugMode"`
}

type FolderStructure struct {
	Resources string `json:"resources"`
	Data      string `json:"data"`
	Packets   string `json:"packets"`
	Scripts   string `json:"scripts"`
	Plugins   string `json:"plugins"`
	Cache     string `json:"cache"`
}

type ServerParam struct {
	ConnectionUri string `json:"connectionUri"`
	Collection    string `json:"collection"`
}

type DatabaseInfo struct {
	Server *ServerParam `json:"server"`
	Game   *ServerParam `json:"game"`
}

type Language struct {
	Language string `json:"language"`
	Fallback string `json:"fallback"`
	Document string `json:"document"`
}

type Account struct {
	AutoCreate          bool     `json:"autoCreate"`
	ExperimentalRealPwd bool     `json:"EXPERIMENTAL_RealPassword"`
	DefaultPermissions  []string `json:"defaultPermissions"`
	MaxPlayer           int      `json:"maxPlayer"`
}

type Http struct {
	StartImmediately bool   `json:"startImmediately"`
	BindAddress      string `json:"bindAddress"`
	BindPort         int    `json:"bindPort"`
	AccessAddress    string `json:"accessAddress"`
	AccessPort       int    `json:"accessPort"`
	Encryption       struct {
		UseEncryption    bool   `json:"useEncryption"`
		UseInRouting     bool   `json:"useInRouting"`
		Keystore         string `json:"keystore"`
		KeystorePassword string `json:"keystorePassword"`
	} `json:"encryption"`
	Policies struct {
		Cors struct {
			Enabled        bool     `json:"enabled"`
			AllowedOrigins []string `json:"allowedOrigins"`
		} `json:"cors"`
	} `json:"policies"`
	Files struct {
		IndexFile string `json:"indexFile"`
		ErrorFile string `json:"errorFile"`
	} `json:"files"`
}

type Game struct {
	BindAddress                string `json:"bindAddress"`
	BindPort                   int    `json:"bindPort"`
	AccessAddress              string `json:"accessAddress"`
	AccessPort                 int    `json:"accessPort"`
	UseUniquePacketKey         bool   `json:"useUniquePacketKey"`
	LoadEntitiesForPlayerRange int    `json:"loadEntitiesForPlayerRange"`
	EnableScriptInBigWorld     bool   `json:"enableScriptInBigWorld"`
	EnableConsole              bool   `json:"enableConsole"`
	KcpInterval                int    `json:"kcpInterval"`
	LogPackets                 string `json:"logPackets"`
	IsShowPacketPayload        bool   `json:"isShowPacketPayload"`
	IsShowLoopPackets          bool   `json:"isShowLoopPackets"`
	CacheSceneEntitiesEveryRun bool   `json:"cacheSceneEntitiesEveryRun"`
	GameOptions                struct {
		InventoryLimits struct {
			Weapons   int `json:"weapons"`
			Relics    int `json:"relics"`
			Materials int `json:"materials"`
			Furniture int `json:"furniture"`
			All       int `json:"all"`
		} `json:"inventoryLimits"`
		AvatarLimits struct {
			SinglePlayerTeam int `json:"singlePlayerTeam"`
			MultiplayerTeam  int `json:"multiplayerTeam"`
		} `json:"avatarLimits"`
		SceneEntityLimit int  `json:"sceneEntityLimit"`
		WatchGachaConfig bool `json:"watchGachaConfig"`
		EnableShopItems  bool `json:"enableShopItems"`
		StaminaUsage     bool `json:"staminaUsage"`
		EnergyUsage      bool `json:"energyUsage"`
		FishhookTeleport bool `json:"fishhookTeleport"`
		TrialCostumes    bool `json:"trialCostumes"`
		Questing         struct {
			Enabled bool `json:"enabled"`
		} `json:"questing"`
		ResinOptions struct {
			ResinUsage   bool `json:"resinUsage"`
			Cap          int  `json:"cap"`
			RechargeTime int  `json:"rechargeTime"`
		} `json:"resinOptions"`
		Rates struct {
			AdventureExp float64 `json:"adventureExp"`
			Mora         float64 `json:"mora"`
			LeyLines     float64 `json:"leyLines"`
		} `json:"rates"`
		Handbook struct {
			Enable        bool `json:"enable"`
			AllowCommands bool `json:"allowCommands"`
			Limits        struct {
				Enabled     bool `json:"enabled"`
				Interval    int  `json:"interval"`
				MaxRequests int  `json:"maxRequests"`
				MaxEntities int  `json:"maxEntities"`
			} `json:"limits"`
			Server struct {
				Enforced  bool   `json:"enforced"`
				Address   string `json:"address"`
				Port      int    `json:"port"`
				CanChange bool   `json:"canChange"`
			} `json:"server"`
		} `json:"handbook"`
	} `json:"gameOptions"`
	JoinOptions struct {
		WelcomeEmotes  []int  `json:"welcomeEmotes"`
		WelcomeMessage string `json:"welcomeMessage"`
		WelcomeMail    struct {
			Title   string `json:"title"`
			Content string `json:"content"`
			Sender  string `json:"sender"`
			Items   []struct {
				ItemId    int `json:"itemId"`
				ItemCount int `json:"itemCount"`
				ItemLevel int `json:"itemLevel"`
			} `json:"items"`
		} `json:"welcomeMail"`
	} `json:"joinOptions"`
	ServerAccount struct {
		AvatarId      int    `json:"avatarId"`
		NameCardId    int    `json:"nameCardId"`
		AdventureRank int    `json:"adventureRank"`
		WorldLevel    int    `json:"worldLevel"`
		NickName      string `json:"nickName"`
		Signature     string `json:"signature"`
	} `json:"serverAccount"`
	VisionOptions []struct {
		Name        string `json:"name"`
		VisionRange int    `json:"visionRange"`
		GridWidth   int    `json:"gridWidth"`
	} `json:"visionOptions"`
}

type Dispatch struct {
	Regions       []*Region `json:"regions"`
	DispatchUrl   string    `json:"dispatchUrl"`
	EncryptionKey string    `json:"encryptionKey"`
	DispatchKey   string    `json:"dispatchKey"`
	DefaultName   string    `json:"defaultName"`
	LogRequests   string    `json:"logRequests"`
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
