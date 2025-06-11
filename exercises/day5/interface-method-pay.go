// create method and interface payment
package main

import "fmt"

type Payer interface {
	pay(amount int)
}

type (
	Cash   struct{}
	Wallet struct{}
	Card   struct{}
)

func (c Cash) pay(amount int) {
	fmt.Printf("paid %d with cash\n", amount)
}

func (w Wallet) pay(amount int) {
	fmt.Printf("paid %d with Wallet\n", amount)
}

func (c Card) pay(amount int) {
	fmt.Printf("paid %d with Card\n", amount)
}

func Dopayment(p Payer, amount int) {
	p.pay(amount)
}

func main() {
	var amount int
	var choice int
	fmt.Println("please enter amount")
	fmt.Scan(&amount)
	fmt.Println("1. Cash")
	fmt.Println("2. Wallet")
	fmt.Println("3. Card")
	fmt.Println("choose an option")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		Dopayment(Cash{}, amount)
	case 2:
		Dopayment(Wallet{}, amount)
	case 3:
		Dopayment(Card{}, amount)
	}
}
