package facade

import "fmt"

type Wallet struct {
	Balance int
}

func NewWallet() *Wallet {
	return &Wallet{
		Balance: 0,
	}
}

func (w *Wallet) CreditBalance(amount int) {
	w.Balance += amount
	fmt.Println("Wallet balance added successfully")
	return
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.Balance < amount {
		return fmt.Errorf("Balance is not sufficient\n")
	}
	fmt.Println("Wallet balance is Sufficient")
	w.Balance -= amount
	return nil
}
