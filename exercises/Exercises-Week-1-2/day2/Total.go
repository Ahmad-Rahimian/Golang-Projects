package main

import "fmt"

func totalDefference(number1, number2 int) (int, int) {
	sum := number1 + number2
	difference := number1 - number2
	return sum, difference
}

func main() {
	var number1, number2 int
	fmt.Println("Please enter first number : ")
	fmt.Scan(&number1)
	fmt.Println("Please enter second number : ")
	fmt.Scan(&number2)
	sum, difference := totalDefference(number1, number2)

	fmt.Printf("Sum is: %d\n", sum)
	fmt.Printf("Difference is: %d\n", difference)
}
