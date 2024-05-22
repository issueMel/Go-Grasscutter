package player

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
