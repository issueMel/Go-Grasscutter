package world

import "Go-Grasscutter/generated/pb"

type Position struct {
	X float32 `bson:"x" json:"X"`
	Y float32 `bson:"y" json:"Y"`
	Z float32 `bson:"z" json:"Z"`
}

var (
	serialVersionUID int64 = -2001232313615923575
	ZERO                   = Position{0, 0, 0}
	IDENTITY               = Position{0, 0, 0}
)

func (p *Position) ToProto() *pb.Vector {
	return &pb.Vector{
		X: p.X,
		Y: p.Y,
		Z: p.Z,
	}
}
