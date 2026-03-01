package account

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"sync"
)

// Define the Account type here.
type Account struct {
	Number string
	Bal    int64
	Open   bool
	Mu     sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	max := big.NewInt(1000)

	n, err := rand.Int(rand.Reader, max)

	if err != nil {
		panic(err)
	}

	b := make([]byte, n.Int64())

	if _, err := rand.Read(b); err != nil {
		return nil
	}

	// Use URL-safe encoding to make it suitable for use in URLs/filenames
	accNum := base64.URLEncoding.EncodeToString(b)

	return &Account{
		Number: accNum,
		Bal:    amount,
		Open:   true,
	}
}

func (a *Account) Balance() (int64, bool) {
	a.Mu.Lock()
	defer a.Mu.Unlock()

	if !a.Open {
		return 0, false
	}

	return a.Bal, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Mu.Lock()
	defer a.Mu.Unlock()

	if !a.Open {
		return 0, false
	}

	if amount < 0 && a.Bal < -amount {
		return 0, false
	}

	a.Bal += amount

	return a.Bal, true
}

func (a *Account) Close() (int64, bool) {
	a.Mu.Lock()
	defer a.Mu.Unlock()

	if !a.Open {
		return 0, false
	}

	a.Open = false
	bal := a.Bal
	a.Bal = 0

	return bal, true
}
