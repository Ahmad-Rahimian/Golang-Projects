package main

import "fmt"

func addnumber() {
	var i, number, n int
	var numbers []int
	fmt.Println("please enter how many number?")
	fmt.Scan(&number)
	for i = 0; i < number; i++ {
		fmt.Println("please enter number :")
		fmt.Scan(&n)
		numbers = append(numbers, n)
	}
	fmt.Println("numbers list :", numbers)
}

func main() {
	addnumber()
}
