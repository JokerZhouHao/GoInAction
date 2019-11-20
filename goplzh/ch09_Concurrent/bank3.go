package main

import "sync"

var(
	mu sync.Mutex
	balance int
)

func Deposit3(amount int)  {
	mu.Lock()
	defer mu.Unlock()
	balacne = balacne + amount
}

func Balance3() int {
	mu.Lock()
	defer mu.Unlock()
	b := balacne
	mu.Unlock()
	return b
}

func deposit3(amount int)  {
	balacne += amount
}

func Withdraw3(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit3(-amount)
	if balance < 0 {
		deposit3(amount)
		return false
	}
	return true
}

var mu1 sync.RWMutex

func Balance33() int {
	mu1.RLock()
	defer mu1.RUnlock()
	return balacne
}
