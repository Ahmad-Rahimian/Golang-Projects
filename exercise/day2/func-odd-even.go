package main

import "fmt"

func oddeven(number int) string {
	if number%2 == 0 {
		return fmt.Sprintf("%v is even", number)
	} else {
		return fmt.Sprintf("%v is odd\n", number)
	}
}

func main() {
	var number int
	fmt.Println("Please enter number : ")
	fmt.Scan(&number)
	result := oddeven(number)
	fmt.Println(result)
}
