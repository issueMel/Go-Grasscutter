package dispatch

import (
	"Go-Grasscutter/src/server/event/types"
)

// todo modify struct
type QueryCurrentRegionEvent struct {
	types.ServerEvent
	RegionInfo string
}

func NewQueryCurrentRegionEvent(regionInfo string) *QueryCurrentRegionEvent {
	return &QueryCurrentRegionEvent{
		// todo modify struct
		ServerEvent: types.ServerEvent{
			Type: types.DISPATCH,
		},
		RegionInfo: regionInfo,
	}
}
