package main

import "fmt"

func isprime(number int) bool {
	if number < 2 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number int
	fmt.Println("Please enter number : ")
	fmt.Scan(&number)
	if isprime(number) {
		fmt.Printf("%v is a prime number\n", number)
	} else {
		fmt.Printf("%v is not a prime number\n", number)
	}
}
