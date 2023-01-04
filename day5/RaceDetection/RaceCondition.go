package main

import (
	"fmt"
	"sync"
)

var AccountBalance int = 5000
var wg1 sync.WaitGroup
var mutex sync.Mutex

func main() {

	wg1.Add(2)
	go deposit(100000)
	go viewBalance()
	wg1.Wait()
}

func deposit(amount int) {
	mutex.Lock()
	AccountBalance = AccountBalance + amount
	mutex.Unlock()
	wg1.Done()
}

func viewBalance() {
	mutex.Lock()
	fmt.Printf("Account Balance:%d", AccountBalance)
	mutex.Unlock()
	wg1.Done()
}
