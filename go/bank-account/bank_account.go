// Package account provides a basic bank account implemetation.
package account

import "sync"

// Account represents a bank account.
type Account struct {
	balance int64
	closed  bool
	sync.Mutex
}

// Open opens a new bank account.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	a := Account{
		balance: initialDeposit,
		closed:  false,
	}

	return &a
}

// Close closes a bank account, returning the remaining balance and whether
// or not the close operation was successful.
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	payout = a.balance
	a.balance = 0
	a.closed = true

	return payout, true
}

// Balance returns the balance for a bank account along with whether or not the
// bank account is closed.
func (a *Account) Balance() (balance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	return a.balance, !a.closed
}

// Deposit updates a bank account balance with a deposit or withdrawal,
// returning the new balance and whether or not the operation was successful.
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if a.closed {
		return 0, false
	}

	a.Lock()
	defer a.Unlock()

	if a.balance+amount < 0 {
		return a.balance, false
	}

	a.balance += amount

	return a.balance, true
}
