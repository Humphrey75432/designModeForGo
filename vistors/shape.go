package vistors

type Shape interface {
	GetType() string
	Accept(visitor Visitor)
}
