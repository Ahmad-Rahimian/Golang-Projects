// Summary: Implements transfer of money between user-defined bank accounts with interactive user input.

package main

import "fmt"

type BankAccount interface {
	Withdraw(amount int) bool
	Deposit(amount int)
	GetBalance() int
}

// =================== Structs ===================

type SavingAccount struct {
	Balance int
}
type CheckingAccount struct {
	Balance int
}

// =================== SavingAccount methods ===================

func (s *SavingAccount) Withdraw(amount int) bool {
	if s.Balance >= amount {
		s.Balance -= amount
		return true
	}
	return false
}

func (s *SavingAccount) Deposit(amount int) {
	s.Balance += amount
}

func (s *SavingAccount) GetBalance() int {
	return s.Balance
}

// =================== CheckingAccount methods ===================

func (c *CheckingAccount) Withdraw(amount int) bool {
	if c.Balance >= amount {
		c.Balance -= amount
		return true
	}
	return false
}

func (c *CheckingAccount) Deposit(amount int) {
	c.Balance += amount
}

func (c *CheckingAccount) GetBalance() int {
	return c.Balance
}

// =================== Transfer Function ===================

func transfer(from BankAccount, to BankAccount, amount int) {
	if from.Withdraw(amount) {
		to.Deposit(amount)
		fmt.Println("✅ Transfer successful")
	} else {
		fmt.Println("❌ Transfer failed: insufficient funds")
	}
}

// =================== Main Function ===================

func main() {
	var savingBalance, checkingBalance int
	var fromChoice, toChoice, amount int

	fmt.Println("Enter initial balance for Saving Account:")
	fmt.Scan(&savingBalance)
	fmt.Println("Enter initial balance for Checking Account:")
	fmt.Scan(&checkingBalance)

	saving := &SavingAccount{Balance: savingBalance}
	checking := &CheckingAccount{Balance: checkingBalance}

	fmt.Println("\nFrom which account do you want to transfer?")
	fmt.Println("1. Saving Account")
	fmt.Println("2. Checking Account")
	fmt.Print("Your choice: ")
	fmt.Scan(&fromChoice)

	fmt.Println("\nTo which account do you want to transfer?")
	fmt.Println("1. Saving Account")
	fmt.Println("2. Checking Account")
	fmt.Print("Your choice: ")
	fmt.Scan(&toChoice)

	fmt.Print("\nEnter amount to transfer: ")
	fmt.Scan(&amount)

	var fromAccount, toAccount BankAccount

	if fromChoice == 1 {
		fromAccount = saving
	} else {
		fromAccount = checking
	}

	if toChoice == 1 {
		toAccount = saving
	} else {
		toAccount = checking
	}

	transfer(fromAccount, toAccount, amount)

	fmt.Println("\n💰 Final Balances:")
	fmt.Println("Saving Account:", saving.GetBalance())
	fmt.Println("Checking Account:", checking.GetBalance())
}
