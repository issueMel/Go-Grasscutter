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

func (b *BargainRecord) ToSnapshot() *pb.BargainSnapshot {
	return &pb.BargainSnapshot{
		BargainId:   uint32(b.BargainId),
		CurMood:     int32(b.CurrentMood),
		BALOPACHCDB: uint32(b.ExpectedPrice),
		IOCNPJJNHLD: uint32(b.LowestPrice),
	}
}
