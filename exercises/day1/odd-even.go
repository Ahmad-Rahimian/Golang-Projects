package main

import "fmt"

func main() {
	var n int
	fmt.Println("please enter number : ")
	fmt.Scan(&n)
	if (n % 2) == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
}
