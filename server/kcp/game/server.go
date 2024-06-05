package game

import (
	"Go-Grasscutter/game/player"
	"sync"
	"sync/atomic"
)

type GameServer struct {
	// todo CHECK: use RWMap or sync.Map
	Players   sync.Map // id, player
	PlayerNum atomic.Int32
}

var Server = &GameServer{
	Players: sync.Map{},
}

func (s *GameServer) RegisterPlayer(player *player.Player) {
	s.Players.Store(player.ID, player)
	s.PlayerNum.Add(1)
}

func (s *GameServer) KickPlayer(id int) {
	s.Players.Delete(id)
	s.PlayerNum.Add(-1)
}
