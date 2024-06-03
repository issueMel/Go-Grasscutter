package excels

import "Go-Grasscutter/generated/pb"

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

func (b *BattlePassMissionData) ToProto() *pb.BattlePassMission {
	prot := &pb.BattlePassMission{
		MissionId:             uint32(b.Id),
		TotalProgress:         uint32(b.Progress),
		RewardBattlePassPoint: uint32(b.AddPoint),
		MissionStatus:         pb.BattlePassMission_MISSION_STATUS_UNFINISHED,
		MissionType:           0,
	}

	if len(b.RefreshType) != 0 {
		// todo BattlePassMissionRefreshType
	}

	return prot
}
