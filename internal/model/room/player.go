package room

type Player struct {
	Name string
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
	}
}

func (p *Player) GetName() string {
	return p.Name
}
func (p *Player) SetName(name string) {
	p.Name = name
}
