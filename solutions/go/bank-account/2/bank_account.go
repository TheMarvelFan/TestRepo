package bankaccount

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"sync"
)

type Account struct {
	Number string
	Bal    int64
	Open   bool
	mu     sync.Mutex
}

func Open(amt int64) *Account {
	if amt < 0 {
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
		Bal:    amt,
		Open:   true,
	}
}

func (a *Account) Balance() (bal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.Open {
		bal = 0
        ok = false
        return
	}

	bal = a.Bal
    ok = true
    return
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.Open {
		newBal = 0
        ok = false
        return
	}

    if amt < 0 && a.Bal < -amt {
        newBal = 0
        ok = false
        return
    }

	a.Bal += amt

	newBal = a.Bal
    ok = true
    return
}

func (a *Account) Withdraw(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.Open || amt < 0 || a.Bal < amt {
		newBal = 0
        ok = false
        return
	}

	a.Bal -= amt

	newBal = a.Bal
    ok = true
    return
}

func (a *Account) Close() (pay int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if !a.Open {
		pay = 0
        ok = false
        return
	}

	a.Open = false
	bal := a.Bal
	a.Bal = 0

	pay = bal
    ok = true
    return
}
