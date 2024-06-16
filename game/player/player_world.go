package player

import (
	"Go-Grasscutter/game/entity"
	"Go-Grasscutter/game/pros"
	"Go-Grasscutter/game/world"
	"time"
)

func (p *Player) InitWorld() {
	p.World = &world.World{
		WorldLevel:       p.GetWorldLevel(),
		IsMultiplayer:    false,
		TimeLocked:       p.GetProperty(pros.PROP_IS_GAME_TIME_LOCKED) != 0,
		LastUpdateTime:   time.Now().UnixMilli(),
		CurrentWorldTime: p.PlayerGameTime,
		Entity: &world.EntityWorld{
			GameEntity: &entity.GameEntity{
				Id:                  0,
				BlockId:             0,
				ConfigId:            0,
				GroupId:             0,
				MotionState:         0,
				LastMoveSceneTimeMs: 0,
				LastMoveReliableSeq: 0,
				LockHP:              false,
				Limbo:               false,
				LimboHpThreshold:    0,
				IsDead:              false,
			},
		},
	}
}

func (p *Player) GetWorldLevel() int {
	return p.GetProperty(pros.PROP_PLAYER_WORLD_LEVEL)
}
