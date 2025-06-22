package vistors

type Square struct {
	Side int
}

func (s *Square) GetType() string {
	return "Square"
}

func (s *Square) Accept(visitor Visitor) {
	visitor.VisitForSquare(s)
}
