package achievement

type Achievement struct {
	Status             string `bson:"status"` // pb.Achievement_KNKEIELCCDB
	Id                 int    `bson:"id"`
	TotalProgress      int    `bson:"totalProgress"`
	CurProgress        int    `bson:"curProgress"`
	FinishTimestampSec int    `bson:"finishTimestampSec"`
}
