package flyweight

type Player struct {
	Dress      Dress
	PlayerType string
	Lat        int
	Long       int
}

func NewPlayer(playerType, dressType string) *Player {
	return nil
}

func (p *Player) NewLocation(lat, long int) {
	p.Lat = lat
	p.Long = long
}
