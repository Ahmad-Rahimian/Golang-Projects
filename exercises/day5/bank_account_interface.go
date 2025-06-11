// Demonstrates using interfaces to handle different types of bank accounts with custom withdrawal logic.
package main

import "fmt"

type BankAccount interface {
	Withdraw(amount int)
}

type (
	SavingAccount struct {
		Balance int
	}
	CheckingAccount struct {
		Balance int
	}
	LockedAccount struct{}
)

func (s *SavingAccount) Withdraw(amount int) {
	if s.Balance >= amount {
		s.Balance -= amount
		fmt.Printf("Remaining: %d\n", s.Balance)
	} else {
		fmt.Println("Not enough balance")
	}
}

func (c *CheckingAccount) Withdraw(amount int) {
	c.Balance -= amount
	fmt.Printf("Remaining: %d\n", c.Balance)
}

func (l LockedAccount) Withdraw(amount int) {
	fmt.Println("Access denied: Locked account")
}

func ProcessWithdrawal(account BankAccount, amount int) {
	account.Withdraw(amount)
}

func main() {
	var amount, choice int
	fmt.Println("please enter amount")
	fmt.Scan(&amount)

	saving := &SavingAccount{Balance: 500}
	checking := &CheckingAccount{Balance: 500}
	locked := LockedAccount{}

	fmt.Println("1. savingaccount")
	fmt.Println("2. checkingaccount")
	fmt.Println("3. lockedaccount")
	fmt.Println("choose an option")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		ProcessWithdrawal(saving, amount)
	case 2:
		ProcessWithdrawal(checking, amount)
	case 3:
		ProcessWithdrawal(locked, amount)
	}
}
