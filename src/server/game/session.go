package game

import (
	"Go-Grasscutter/src/config"
	"Go-Grasscutter/src/game"
	"Go-Grasscutter/src/game/player"
	"Go-Grasscutter/src/utils/crypto"
	"bytes"
	"encoding/binary"
	"log"
	"time"
)

type GameSession struct {
	Server *GameServer
	// Tunnel
	Account       *game.Account
	Player        *player.Player
	EncryptSeed   uint64
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

func handlerReceive(buffer []byte) {
	// todo use EncryptKey / DispatchKey
	buf := bytes.NewBuffer(crypto.Xor(buffer, crypto.DispatchKey))
	for buf.Len() > 0 {
		if buf.Len() < 12 {
			break
		}
		var const1, const2 int16
		var opcode, headerLength uint16
		var payloadLength uint32
		err := binary.Read(buf, binary.BigEndian, &const1)
		if err != nil {
			log.Println("Read const1 error:", err)
			break
		}
		// Packet sanity check
		if const1 != 17767 {
			// Bad packet
			log.Printf("Bad Data Package Received: got %d ,expect 17767", const1)
			break
		}

		// Data
		opcode = binary.BigEndian.Uint16(buf.Next(2))
		headerLength = binary.BigEndian.Uint16(buf.Next(2))
		payloadLength = binary.BigEndian.Uint32(buf.Next(4))
		header := make([]byte, headerLength)
		payload := make([]byte, payloadLength)

		_, err = buf.Read(header)
		if err != nil {
			log.Println("Read header error:", err)
			break
		}
		_, err = buf.Read(payload)
		if err != nil {
			log.Println("Read payload error:", err)
			break
		}

		// Sanity check #2
		err = binary.Read(buf, binary.BigEndian, &const2)
		if err != nil {
			log.Println("Read short const2 error:", err)
			break
		}
		if const2 != -30293 {
			log.Printf("Bad Data Package Received: got %d ,expect -30293", const2)
			break // Bad packet
		}
		// todo Handle
		log.Println("handle", opcode)
	}
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
