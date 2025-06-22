package states

import "fmt"

type NoItemState struct {
	vendingMachine *VendingMachine
}

func (n *NoItemState) AddItem(i int) error {
	fmt.Printf("%d items added\n", i)
	n.vendingMachine.incrementItemCount(i)
	return nil
}

func (n *NoItemState) RequestItem() error {
	if n.vendingMachine.itemCount == 0 {
		n.vendingMachine.setState(n.vendingMachine.noItem)
		return fmt.Errorf("No item present\n")
	}
	fmt.Printf("Item requested\n")
	n.vendingMachine.setState(n.vendingMachine.itemRequested)
	return nil
}

func (n *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("Please select item first\n")
}

func (n *NoItemState) DispenseItem() error {
	return fmt.Errorf("Please select item first\n")
}
