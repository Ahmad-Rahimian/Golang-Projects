// A goroutine sends a Book struct into a channel, and the main function receives and prints it.
package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Year   int
}

func main() {
	ch := make(chan Book)
	go func() {
		book := Book{Title: "noah", Author: "ahmad", Year: 1990}
		ch <- book
	}()

	bookdetail := <-ch
	fmt.Println("book detail :", bookdetail)
}
