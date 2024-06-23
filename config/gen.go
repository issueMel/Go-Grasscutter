package config

import (
	_ "embed"
	"encoding/json"
	"os"
)

//go:embed config.json
var config []byte

// Deprecated
func genConfigFileByEmbed() {
	filepath := "config.json"
	err := os.WriteFile(filepath, config, 0644)
	if err != nil {
		panic(err)
	}
}

func genConfigFile() {
	genConfig()
	marshal, err := json.MarshalIndent(Conf, "", "	")
	if err != nil {
		panic(err)
	}

	filepath := "config.json"
	err = os.WriteFile(filepath, marshal, 0644)
	if err != nil {
		panic(err)
	}
}

func genConfig() {
	Conf = &Config{}

	// DatabaseInfo
	Conf.DatabaseInfo = NewDatabaseInfo()

	// FolderStructure
	Conf.FolderStructure = NewStructure()

	// Account
	Conf.Account = NewAccount()

	// Server
	Conf.Server = NewServer()

}

func NewDatabaseInfo() *DatabaseInfo {
	datainfo := &DataStore{
		ConnectionUri: "mongodb://localhost:27017",
		Collection:    "grasscutter",
	}
	return &DatabaseInfo{
		Server: datainfo,
	}
}

func NewStructure() *FolderStructure {
	struc := &FolderStructure{}
	struc.Resources = "./resources/"

	return struc
}

func NewAccount() *Account {
	account := &Account{}
	account.MaxPlayer = -1

	return account
}

func NewServer() *Server {
	s := &Server{}
	s.RunMode = "HYBRID" // HYBRID, DISPATCH_ONLY, GAME_ONLY todo ServerRunMode
	s.Http = NewHTTP()
	s.Game = NewGame()
	s.Dispatch = NewDispatch()
	return s
}

func NewHTTP() *Http {
	h := &Http{}
	h.BindAddress = "0.0.0.0"
	h.BindPort = 443

	h.AccessAddress = "127.0.0.1"
	h.AccessPort = 0

	return h
}

func NewGame() *Game {
	g := &Game{}
	g.BindAddress = "0.0.0.0"
	g.BindPort = 22102
	g.AccessAddress = "127.0.0.1"
	g.AccessPort = 0
	// Enabling this will generate a unique packet encryption key for each player.
	g.UseUniquePacketKey = true

	// Kcp internal work interval (milliseconds)
	g.KcpInterval = 20

	g.GameOptions = NewGameOptions()
	return g
}

func NewDispatch() *Dispatch {
	return &Dispatch{
		Regions: []*Region{
			{
				Name:  "os_usa",
				Title: "Grasscutter",
				Ip:    "127.0.0.1",
				Port:  22102,
			},
		},
		DefaultName: "Grasscutter",
	}
}

func NewGameOptions() *GameOptions {
	gopt := &GameOptions{}
	gopt.InventoryLimits = NewInventoryLimits()
	gopt.ResinOptions = NewReResinOptions()
	return gopt
}

func NewInventoryLimits() *InventoryLimits {
	return &InventoryLimits{
		Weapons:   2000,
		Relics:    2000,
		Materials: 2000,
		Furniture: 2000,
		All:       30000,
	}
}

func NewReResinOptions() *ResinOptions {
	return &ResinOptions{
		ResinUsage:   false,
		Cap:          160,
		RechargeTime: 480,
	}
}
