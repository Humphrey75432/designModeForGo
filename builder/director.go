package builder

type Director struct {
	Builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.Builder = b
}

func (d *Director) BuildHouse() House {
	d.Builder.SetWindowType()
	d.Builder.SetDoorType()
	d.Builder.SetNumFloor()
	return d.Builder.GetHouse()
}
