package medium

import "fmt"

type FreightTrain struct {
	Mediator Mediator
}

func (f *FreightTrain) Arrive() {
	if !f.Mediator.CanArrive(f) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

func (f *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	f.Mediator.NotifyAboutDeparture()
}

func (f *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	f.Arrive()
}
