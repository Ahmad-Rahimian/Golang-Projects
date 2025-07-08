// Summary: Print even and odd numbers concurrently using Goroutines.
package main

import (
	"fmt"
	"time"
)

func printEven() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even", i)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func printOdd() {
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("odd", i)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	go printEven()
	go printOdd()
	time.Sleep(5 * time.Second)
}
