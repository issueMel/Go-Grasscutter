package player

type PlayerBirthday struct {
	Day   int
	Month int
}

func (p *PlayerBirthday) PlayerBirthday() {
	p.Day = 0
	p.Month = 0
}
