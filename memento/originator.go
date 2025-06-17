package memento

type Originator struct {
	State string
}

func (o *Originator) CreateMemento() *Memento {
	return &Memento{State: o.State}
}

func (o *Originator) RestoreMemento(m *Memento) {
	o.State = m.GetSavedState()
}

func (o *Originator) SetState(state string) {
	o.State = state
}

func (o *Originator) GetState() string {
	return o.State
}
