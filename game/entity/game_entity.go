package entity

import "Go-Grasscutter/generated/pb"

type GameEntity struct {
	Id int

	BlockId             int
	ConfigId            int
	GroupId             int
	MotionState         pb.MotionState
	LastMoveSceneTimeMs int
	LastMoveReliableSeq int
	LockHP              bool
	Limbo               bool
	LimboHpThreshold    float32
	IsDead              bool
}
