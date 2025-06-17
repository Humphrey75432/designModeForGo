package medium

type Mediator interface {
	CanArrive(train Train) bool
	NotifyAboutDeparture()
}
