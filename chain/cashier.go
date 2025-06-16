package chain

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(patient *Patient) {
	if patient.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
