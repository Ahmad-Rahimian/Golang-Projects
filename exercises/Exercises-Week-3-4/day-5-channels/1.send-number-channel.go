// A goroutine sends a number into a channel, and the main function receives and prints it.
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	num := <-ch
	fmt.Println("recived :", num)
}
