package excels

type BattlePassMissionData struct {
	Id            int
	AddPoint      int
	ScheduleId    int
	Progress      int
	TriggerConfig TriggerConfig
	RefreshType   string // todo BattlePassMissionRefreshType
	MainParams    map[int]struct{}
}

type TriggerConfig struct {
	TriggerType string // todo WatcherTriggerType
	ParamList   []string
}

func (b *BattlePassMissionData) IsValidRefreshType() bool {
	return b.RefreshType == "" ||
		b.RefreshType == "BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE" ||
		b.ScheduleId == 2701
}
