package builder

type NormalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func (n *NormalBuilder) SetWindowType() {
	n.WindowType = "Wooden Window"
}

func (n *NormalBuilder) SetDoorType() {
	n.DoorType = "Wooden Door"
}

func (n *NormalBuilder) SetNumFloor() {
	n.Floor = 2
}

func (n *NormalBuilder) GetHouse() House {
	return House{
		DoorType:   n.DoorType,
		WindowType: n.WindowType,
		Floor:      n.Floor,
	}
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}
