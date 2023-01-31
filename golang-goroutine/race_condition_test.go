package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	var counter int32 = 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(
		"The counter expected result is",
		(1000 * 100),
		"but got actual result",
		counter)
}

func TestRaceConditionWithMutex(t *testing.T) {
	var counter int32 = 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println(
		"The counter expected result is",
		(1000 * 100),
		"but got actual result",
		counter)
}

func TestRaceConditionWithRWMutex(t *testing.T) {
	var bankAccount BankAccount = BankAccount{
		Id:      "1234FRDE",
		Balance: 1000.00,
	}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				bankAccount.AddBalance(1)
				fmt.Println(bankAccount.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Total Balance :", bankAccount.GetBalance())
}

func TestDeadlock(t *testing.T) {
	var bca BankAccount = BankAccount{
		Id:      "1234IZYA",
		Balance: 1000.00,
	}
	var bni BankAccount = BankAccount{
		Id:      "9876TFRX",
		Balance: 1000.00,
	}

	go bca.Transfer(&bni, 1000)
	go bni.Transfer(&bca, 1000)

	time.Sleep(5 * time.Second)

	fmt.Println("BCA Balance :", bca.GetBalance())
	fmt.Println("BNI Balance :", bni.GetBalance())

}
