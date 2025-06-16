package facade

import "fmt"

type Account struct {
	Name string
}

func NewAccount(accountName string) *Account {
	return &Account{
		Name: accountName,
	}
}
func (a *Account) CheckAccount(accountName string) error {
	if a.Name != accountName {
		return fmt.Errorf("Account Name %s is incorrect\n", accountName)
	}
	fmt.Println("Account Verified")
	return nil
}
