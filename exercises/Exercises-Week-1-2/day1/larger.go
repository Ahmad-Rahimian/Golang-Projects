package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Println("enter the number : ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error: Please enter a valid integer number.")
		os.Exit(1)
	}
	max := 0
	for {
		b := n % 10
		n /= 10
		if max < b {
			max = b
		}
		if n == 0 {
			break
		}

	}
	fmt.Println("larger : ", max)
}
