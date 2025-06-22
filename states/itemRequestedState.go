package states

import "fmt"

type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("Item Dispense in progress\n")
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("Item already requested\n")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("Insert money is less. Please insert %d\n", i.vendingMachine.itemPrice)
	}
	fmt.Printf("Money entered is ok\n")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("Please inser money first\n")
}
