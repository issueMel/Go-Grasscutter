package player

import "Go-Grasscutter/generated/pb"

type PlayerBirthday struct {
	Day   int
	Month int
}

func (p *PlayerBirthday) NewPlayerBirthday() {
	p.Day = 0
	p.Month = 0
}

func (p *PlayerBirthday) SetPlayerBirthday(day, month int) {
	p.Day = day
	p.Month = month
}

func (p *PlayerBirthday) GetFilledProtoWhenNotEmpty() *pb.Birthday {
	if p.Day > 0 {
		return &pb.Birthday{
			Month: uint32(p.Month),
			Day:   uint32(p.Day),
		}
	}
	return &pb.Birthday{}
}
