package main

import "fmt"

func addNumbers() {
	var i, number, n, average int
	var numbers []int
	var sum int
	fmt.Println("please enter how many numbers?")
	fmt.Scan(&number)
	for i = 0; i < number; i++ {
		fmt.Printf("please enter number %d:\n", i+1)
		fmt.Scan(&n)
		numbers = append(numbers, n)
	}
	for _, n := range numbers {
		sum += n
		average = sum / number
	}
	fmt.Printf("average numbers : %v\n", average)
}

func main() {
	addNumbers()
}
