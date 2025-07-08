// Summary: Simulates multiple clients sending messages concurrently using Goroutines.
package main

import (
	"fmt"
	"time"
)

func client(id int) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("Client %d sends message %d\n", id, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	for i := 1; i <= 5; i++ {
		go client(i)
	}

	time.Sleep(3 * time.Second)
}
