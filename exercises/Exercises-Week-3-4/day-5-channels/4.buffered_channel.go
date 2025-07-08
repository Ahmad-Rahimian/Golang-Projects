// Summary: Demonstrates how a buffered channel can hold multiple values before they're received.
package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 3)

	messages <- "message 1"
	messages <- "message 2"
	messages <- "message 3"

	fmt.Println("Length:", len(messages))
	fmt.Println("Capacity:", cap(messages))

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
