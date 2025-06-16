package chain

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(patient *Patient) {
	if patient.MedicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(patient)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	patient.MedicineDone = true
	m.next.Execute(patient)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}
