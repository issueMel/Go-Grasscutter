package types

import "Go-Grasscutter/src/server/event"

type Type int

type ServerEvent struct {
	Event event.Event
	Type  Type
}

const (
	DISPATCH Type = iota
	GAME
)
