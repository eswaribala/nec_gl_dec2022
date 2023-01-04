package main

import (
	"fmt"
	"sync"
)

var AccountBalance int = 5000
var wg1 sync.WaitGroup

func main() {

	wg1.Add(2)
	go deposit(100000)
	go viewBalance()
	wg1.Wait()
}

func deposit(amount int) {

	AccountBalance = AccountBalance + amount

	wg1.Done()
}

func viewBalance() {

	fmt.Printf("Account Balance:%d", AccountBalance)
	wg1.Done()
}
