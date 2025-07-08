// A goroutine sends a number into a channel, and the main function receives and prints it.
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "hello from goroutine 1"
	}()

	go func() {
		ch2 <- "hello from goroutine 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("recived :", msg1)
	case msg2 := <-ch2:
		fmt.Println("recived : ", msg2)

	}
}
