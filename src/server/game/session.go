package game

import (
	"Go-Grasscutter/src/config"
	"Go-Grasscutter/src/game"
	"Go-Grasscutter/src/game/player"
	"Go-Grasscutter/src/utils/crypto"
	"time"
)

type GameSession struct {
	Server *GameServer
	// Tunnel
	Account       *game.Account
	Player        *player.Player
	EncryptSeed   int64
	EncryptKey    []byte
	UseSecretKey  bool
	State         SessionState
	ClientTime    int
	LastPingTime  int64
	LastClientSeq int
}

func NewGameSession(server *GameServer) *GameSession {
	g := &GameSession{
		Server:        server,
		EncryptSeed:   crypto.EncryptSeed,
		EncryptKey:    crypto.EncryptKey,
		State:         WaitingForToken,
		LastPingTime:  time.Now().UnixMilli(),
		LastClientSeq: 10,
	}
	if config.Conf.Server.Game.UseUniquePacketKey {
		g.EncryptKey = make([]byte, 4096)
		g.EncryptSeed = crypto.GenerateEncryptKeyAndSeed(g.EncryptKey)
	}
	return g
}

func (s *GameSession) SetPlayer(p *player.Player) {
	s.Player = p
	s.Player.Session = s
	s.Player.Account = s.Account
}

func (s *GameSession) isLoggedIn() bool {
	return s.Player != nil
}

func (s *GameSession) UpdateLastPingTime(clientTime int) {
	s.ClientTime = clientTime
	s.LastPingTime = time.Now().UnixMilli()
}

func (s *GameSession) GetNextClientSequence() int {
	s.LastClientSeq++
	return s.LastClientSeq
}

func replayPacket() {

}

type SessionState int

const (
	Inactive SessionState = iota
	WaitingForToken
	WaitingForLogin
	PickingCharacter
	Active
	AccountBanned
)
