package world

import (
	"Go-Grasscutter/game/entity"
)

type World struct {
	// players []atomic.Pointer[any]

	Entity           *EntityWorld
	NextEntityId     int
	NextPeerId       int
	WorldLevel       int
	IsMultiplayer    bool
	TimeLocked       bool
	LastUpdateTime   int64
	TickCount        int
	IsPaused         bool
	CurrentWorldTime int64
}

type EntityWorld struct {
	*entity.GameEntity
}
