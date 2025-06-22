package vistors

type Circle struct {
	Radius int
}

func (c *Circle) GetType() string {
	return "Circle"
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitForCircle(c)
}
