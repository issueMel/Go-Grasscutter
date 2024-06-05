package player

import "Go-Grasscutter/generated/pb"

type TeamInfo struct {
	Name    string `bson:"name"`
	Avatars []int  `bson:"avatars"`
}

func (t *TeamInfo) ToProto() *pb.AvatarTeam {
	// todo INCOMPLETE: TeamInfo ToProto()
	return &pb.AvatarTeam{}
}
