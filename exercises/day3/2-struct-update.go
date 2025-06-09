// create and update struct and get info
package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}

// add book
func addBook(title string, author string, price float64) Book {
	b := Book{Title: title, Author: author, Price: price}
	return b
}

// update price
func updatePrice(b *Book, newPrice float64) {
	b.Price = newPrice
}

func main() {
	var title, author string
	var price, newPrice float64
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

	// get new price and update and get new info
	fmt.Println("Please enter the new price:")
	fmt.Scan(&newPrice)
	updatePrice(&book, newPrice)
	fmt.Println("✅ Book Info:")
	fmt.Println("Title:", book.Title)
	fmt.Println("Author:", book.Author)
	fmt.Println("Price:", book.Price)
}
