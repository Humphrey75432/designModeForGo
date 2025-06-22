package vistors

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (m *MiddleCoordinates) VisitForSquare(square *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (m *MiddleCoordinates) VisitForCircle(circle *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (m *MiddleCoordinates) VisitForRectangle(rectangle *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
