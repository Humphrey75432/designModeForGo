package vistors

type Visitor interface {
	VisitForSquare(square *Square)
	VisitForCircle(circle *Circle)
	VisitForRectangle(rectangle *Rectangle)
}
