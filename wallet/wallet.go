package wallet

import (
	"fmt"
	"errors"
)

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(b Bitcoin) {
	w.balance += b
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Withdraw(b Bitcoin) error {
	if b > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= b
	return nil
}
