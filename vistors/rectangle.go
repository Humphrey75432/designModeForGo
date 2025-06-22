package vistors

type Rectangle struct {
	L int
	B int
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitForRectangle(r)
}
