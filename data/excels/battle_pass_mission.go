package excels

import (
	"Go-Grasscutter/config"
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"encoding/json"
	"os"
	"path/filepath"
)

const BattlePassMissionResource = "BattlePassMissionExcelConfigData.json"

type BattlePassMissionData struct {
	Id            int           `json:"id"`
	AddPoint      int           `json:"addPoint"`
	ScheduleId    int           `json:"scheduleId"`
	Progress      int           `json:"progress"`
	TriggerConfig TriggerConfig `json:"triggerConfig"`
	RefreshType   string        `json:"refreshType"` // todo INCOMPLETE: BattlePassMissionRefreshType
	MainParams    map[int]struct{}
}

type TriggerConfig struct {
	TriggerType string   `json:"triggerType"` // todo INCOMPLETE: WatcherTriggerType
	ParamList   []string `json:"paramList"`
}

func (b *BattlePassMissionData) IsValidRefreshType() bool {
	return b.RefreshType == "" ||
		b.RefreshType == "BATTLE_PASS_MISSION_REFRESH_CYCLE_CROSS_SCHEDULE" || // todo INCOMPLETE: BattlePassMissionRefreshType
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
		// todo INCOMPLETE: BattlePassMissionRefreshType
	}

	return prot
}

func LoadBattlePassMissionData() map[int]*BattlePassMissionData {
	prefix := config.Conf.FolderStructure.Resources

	data := []*BattlePassMissionData{}
	m := make(map[int]*BattlePassMissionData)

	fp := filepath.Join(prefix + "ExcelBinOutput/" + BattlePassMissionResource)
	file, err := os.ReadFile(fp)
	if err != nil {
		log.SugaredLogger.Panic(err)
		return nil
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.SugaredLogger.Panic(err)
		return nil
	}

	for _, v := range data {
		m[v.Id] = v
	}
	return m
}
