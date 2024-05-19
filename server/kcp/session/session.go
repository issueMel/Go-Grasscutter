package session

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/game/player"
	"Go-Grasscutter/lib/kcp-go"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/utils/crypto"
	"time"
)

type Session struct {
	Tunnel        *kcp.UDPSession
	Account       *player.Account
	Player        *player.Player
	EncryptSeed   uint64
	EncryptKey    []byte
	UseSecretKey  bool
	State         SessionState
	ClientTime    int
	LastPingTime  int64
	LastClientSeq int
}

func (s *Session) Send(packet *base.BasePacket) {
	if packet.Opcode <= 0 {
		log.SugaredLogger.Warn("Tried to send packet with missing cmd id!")
		return
	}

	// Header
	if packet.ShouldBuildHeader {
		packet.BuildHeader(s.LastClientSeq)
		s.LastClientSeq++
	}

	bytes := packet.Build()

	if packet.ShouldEncrypt {
		if packet.UseDispatchKey {
			crypto.Xor(bytes, crypto.DispatchKey)
		} else {
			crypto.Xor(bytes, s.EncryptKey)
		}
	}

	_, err := s.Tunnel.Write(bytes)
	if err != nil {
		log.SugaredLogger.Debug("Unable to send packet to client:", err)
		return
	}
}

func NewGameSession() *Session {
	g := &Session{
		//Server:        server,
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

func (s *Session) UpdateLastPingTime(clientTime int) {
	s.ClientTime = clientTime
	s.LastPingTime = time.Now().UnixMilli()
}

func (s *Session) GetNextClientSequence() int {
	s.LastClientSeq++
	return s.LastClientSeq
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
