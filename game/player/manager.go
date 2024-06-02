package player

type BaseManager struct {
	Player *Player
}

func (b *BaseManager) SetPlayer(p *Player) {
	b.Player = p
}

func (b *BaseManager) GetPlayer() *Player {
	return b.Player
}

func (b *BaseManager) Save() {
	b.GetPlayer().Save()
}
