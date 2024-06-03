package battlepass

import "Go-Grasscutter/generated/pb"

type Mission struct {
	Id       int    `bson:"id"`
	Progress int    `bson:"progress"`
	Status   string `bson:"status"` // get from pb.BattlePassMission_MissionStatus_value
}

func (m *Mission) ToProto() *pb.BattlePassMission {
	return &pb.BattlePassMission{
		MissionId:   uint32(m.Id),
		CurProgress: uint32(m.Progress),
		//TotalProgress: , // todo get from manager
		MissionStatus: pb.BattlePassMission_MissionStatus(pb.BattlePassMission_MissionStatus_value[m.Status]),
		// MissionType: , // todo get from manager
	}
}
