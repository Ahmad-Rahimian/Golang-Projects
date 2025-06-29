// create struct and get info
package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}

// add book
func addBook(title string, author string, price float64) Book {
	return Book{Title: title, Author: author, Price: price}
}

func main() {
	var title, author string
	var price float64
	fmt.Println("Please enter the book title:")
	fmt.Scan(&title)

	fmt.Println("Please enter the author:")
	fmt.Scan(&author)

	fmt.Println("Please enter the price:")
	fmt.Scan(&price)

	// add book
	book := addBook(title, author, price)

	// get information
	fmt.Println("✅ Book Info:")
	fmt.Println("Title:", book.Title)
	fmt.Println("Author:", book.Author)
	fmt.Println("Price:", book.Price)
}
