package account

import (
	"math"
	"sync"
)

// Define the Account type here.
type Account struct {
	balance int64
	closed  bool
	mu      *sync.Mutex
}

func Open(amount int64) *Account {
	if amount >= 0 {
		return &Account{
			balance: amount,
			closed:  false,
			mu:      &sync.Mutex{},
		}
	}

	return nil
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	if amount < 0 && math.Abs(float64(amount)) > float64(a.balance) {
		return 0, false
	}

	a.balance += amount

	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if a.closed {
		return 0, false
	}

	balance := a.balance
	a.balance = 0
	a.closed = true

	return balance, true
}
