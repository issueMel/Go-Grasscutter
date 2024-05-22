package quest

import "Go-Grasscutter/generated/pb"

type BargainRecord struct {
	BargainId     int                   `bson:"bargainId"`
	LowestPrice   int                   `bson:"lowestPrice"`
	ExpectedPrice int                   `bson:"expectedPrice"`
	CurrentMood   int                   `bson:"currentMood"`
	Finished      bool                  `bson:"finished"`
	Result        *pb.BargainResultType `bson:"result"`
}
