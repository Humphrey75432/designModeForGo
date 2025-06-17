package medium

import "fmt"

type PassengerTrain struct {
	Mediator Mediator
}

func (p *PassengerTrain) Arrive() {
	if !p.Mediator.CanArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.Mediator.NotifyAboutDeparture()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.Arrive()
}
