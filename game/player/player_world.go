package player

import (
	"Go-Grasscutter/game/pros"
	"Go-Grasscutter/game/world"
	"time"
)

func (p *Player) InitWorld() *world.World {
	return &world.World{
		WorldLevel:       p.GetWorldLevel(),
		IsMultiplayer:    false,
		TimeLocked:       p.GetProperty(pros.PROP_IS_GAME_TIME_LOCKED) != 0,
		LastUpdateTime:   time.Now().UnixMilli(),
		CurrentWorldTime: p.PlayerGameTime,
	}
}

func (p *Player) GetWorldLevel() int {
	return p.GetProperty(pros.PROP_PLAYER_WORLD_LEVEL)
}
