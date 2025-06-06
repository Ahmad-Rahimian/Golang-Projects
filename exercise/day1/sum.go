package main

import "fmt"

func main() {
	var a, b float64
	fmt.Println("Please enter first number : ")
	fmt.Scan(&a)
	fmt.Println("Please enter second number : ")
	fmt.Scan(&b)
	sum := a + b
	fmt.Printf("sum : %v\n ", sum)
}
