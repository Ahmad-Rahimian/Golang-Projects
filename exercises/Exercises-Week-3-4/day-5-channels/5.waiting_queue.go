// Summary: Simulates a simple waiting queue using a channel where multiple clients wait to be served.
package main

import (
	"fmt"
	"time"
)

func serve(client int, ch chan int) {
	fmt.Printf("Client %d is waiting...\n", client)
	ch <- client
}

func main() {
	queue := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		go serve(i, queue)
		time.Sleep(200 * time.Millisecond)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("\n--- Serving Clients ---")
	for i := 0; i < 5; i++ {
		client := <-queue
		fmt.Printf("Serving client %d\n", client)
		time.Sleep(500 * time.Millisecond)
	}
}
