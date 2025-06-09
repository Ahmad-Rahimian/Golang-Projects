package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {
	// Create map
	books := make(map[string]Book)

	var count int
	fmt.Print("How many books do you want to enter? ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		var code, title, author string
		var price float64

		fmt.Printf("\nBook #%d\n", i+1)

		fmt.Print("Book code: ")
		fmt.Scan(&code)

		fmt.Print("Title: ")
		fmt.Scan(&title)

		fmt.Print("Author: ")
		fmt.Scan(&author)

		fmt.Print("Price: ")
		fmt.Scan(&price)

		// Create book struct
		book := Book{
			Title:  title,
			Author: author,
			Price:  price,
		}

		// Save to map
		books[code] = book
	}

	// Search
	var searchCode string
	fmt.Print("\nEnter the book code to search: ")
	fmt.Scan(&searchCode)

	if book, found := books[searchCode]; found {
		fmt.Println("\nBook found:")
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Printf("Price: %.2f\n", book.Price)
	} else {
		fmt.Println("\nBook not found.")
	}
}
