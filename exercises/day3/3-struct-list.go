// insert structs in slice and get info
package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Price  float64
}

func addBook(title string, author string, price float64) Book {
	b := Book{Title: title, Author: author, Price: price}
	return b
}

func main() {
	// add book
	var i, number int
	fmt.Println("Please enter how many numbers :")
	fmt.Scan(&number)

	var books []Book

	for i = 0; i < number; i++ {
		var title, author string
		var price float64
		fmt.Printf("\n📘 Book %d\n", i+1)
		fmt.Print("Title: ")
		fmt.Scan(&title)

		fmt.Print("Author: ")
		fmt.Scan(&author)

		fmt.Print("Price: ")
		fmt.Scan(&price)

		book := addBook(title, author, price)
		books = append(books, book)
	}

	fmt.Println("📚 Book List:")
	for i, b := range books {
		fmt.Printf("%d. Title: %s | Author: %s | Price: %.2f\n", i+1, b.Title, b.Author, b.Price)
	}
}
