// Summary: Demonstrates how to trigger a panic for invalid input and recover gracefully using defer and recover.

package main

import (
	"fmt"
)

func checkPositive(n int) {
	if n < 0 {
		panic("negative number not allowed")
	}
}

func main() {
	var n int
	fmt.Println("enter number: ")
	fmt.Scan(&n)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	checkPositive(n)
}
