package observer

type Observer interface {
	Update(string2 string)
	GetID() string
}
