package facade

import "fmt"

type WalletFacade struct {
	Account      *Account
	Wallet       *Wallet
	SecurityCode *SecurityCode
	Notification *Notification
	Ledger       *Ledger
}

func NewWalletFacade(accountId string, code int) *WalletFacade {
	fmt.Println("Starting create account")
	WalletFacade := &WalletFacade{
		Account:      NewAccount(accountId),
		SecurityCode: NewSecurityCode(code),
		Wallet:       NewWallet(),
		Notification: &Notification{},
		Ledger:       &Ledger{},
	}
	fmt.Println("Account created")
	return WalletFacade
}

func (w *WalletFacade) AddMoneyToWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.Account.CheckAccount(accountId)
	if err != nil {
		return err
	}
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	w.Wallet.CreditBalance(amount)
	w.Notification.SendWalletCreditNotification()
	w.Ledger.MakeEntry(accountId, "credit", amount)
	return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.Account.CheckAccount(accountId)
	if err != nil {
		return err
	}
	err = w.SecurityCode.CheckCode(securityCode)
	if err != nil {
		return err
	}
	err = w.Wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.Notification.SendWalletDebitNotification()
	w.Ledger.MakeEntry(accountId, "debit", amount)
	return nil
}
