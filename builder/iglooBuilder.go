package builder

type IglooBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func (i *IglooBuilder) SetWindowType() {
	i.WindowType = "Snow Window"
}

func (i *IglooBuilder) SetDoorType() {
	i.DoorType = "Snow Door"
}

func (i *IglooBuilder) SetNumFloor() {
	i.Floor = 1
}

func (i *IglooBuilder) GetHouse() House {
	return House{
		WindowType: i.WindowType,
		DoorType:   i.DoorType,
		Floor:      i.Floor,
	}
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}
