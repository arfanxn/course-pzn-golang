package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Main program")
}

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func DisplayNumber(number int64) {
	fmt.Println("Number: ", number)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Id      string
	Balance float64
}

func (account *BankAccount) AddBalance(amount float64) {
	account.RWMutex.Lock()
	account.Balance += amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) SetBalance(amount float64) {
	account.Balance = amount
}

func (account *BankAccount) GetBalance() float64 {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func (account *BankAccount) Lock() {
	account.RWMutex.Lock()
}

func (account *BankAccount) Unlock() {
	account.RWMutex.Unlock()
}

func (account *BankAccount) Transfer(destinationAccount *BankAccount, amount float64) {
	account.Lock()
	fmt.Println("First - Lock Bank Account Id :", account.Id)
	account.SetBalance(account.Balance - amount)

	time.Sleep(1 * time.Second)

	destinationAccount.Lock()
	fmt.Println("Second - Lock Bank Account Id :", account.Id)
	destinationAccount.SetBalance(destinationAccount.Balance + amount)

	time.Sleep(1 * time.Second)

	account.Unlock()
	destinationAccount.Unlock()

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)

		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}
