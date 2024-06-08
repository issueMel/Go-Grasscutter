package achievement

import "Go-Grasscutter/generated/pb"

type Achievement struct {
	Status             string `bson:"status"` // pb.Achievement_KNKEIELCCDB
	Id                 int    `bson:"id"`
	TotalProgress      int    `bson:"totalProgress"`
	CurProgress        int    `bson:"curProgress"`
	FinishTimestampSec int    `bson:"finishTimestampSec"`
}

func (a *Achievement) ToProto() *pb.Achievement {
	return &pb.Achievement{
		Status:          pb.Achievement_KNKEIELCCDB(pb.Achievement_KNKEIELCCDB_value[a.Status]),
		Id:              uint32(a.Id),
		TotalProgress:   uint32(a.TotalProgress),
		CurProgress:     uint32(a.CurProgress),
		FinishTimestamp: uint32(a.FinishTimestampSec),
	}
}
