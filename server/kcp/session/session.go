package session

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/game/player"

	"Go-Grasscutter/game/tunnel"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/utils/crypto"
	"time"
)

type Session struct {
	Tunnel       *tunnel.Tunnel
	Account      *player.Account
	Player       *player.Player
	EncryptSeed  uint64
	UseSecretKey bool
	State        SessionState
	ClientTime   uint32
}

func (s *Session) Send(packet *base.Packet) {
	s.Tunnel.Send(packet)
}

func NewGameSession() *Session {
	g := &Session{
		Tunnel: &tunnel.Tunnel{
			LastClientSeq: 10,
			LastPingTime:  time.Now().UnixMilli(),
			EncryptKey:    crypto.EncryptKey,
		},
		EncryptSeed: crypto.EncryptSeed,
		State:       WaitingForToken,
	}
	if config.Conf.Server.Game.UseUniquePacketKey {
		g.Tunnel.EncryptKey, g.EncryptSeed = crypto.GenerateEncryptKeyAndSeed(make([]byte, 4096))
	}
	return g
}

func (s *Session) UpdateLastPingTime(clientTime uint32) {
	s.ClientTime = clientTime
	s.Tunnel.LastPingTime = time.Now().UnixMilli()
}

func (s *Session) GetNextClientSequence() uint32 {
	s.Tunnel.LastClientSeq++
	return s.Tunnel.LastClientSeq
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
