package main

import (
	"bufio"
	"fmt"
	"os"
)

// create struct for Product
type Product struct {
	ID    int
	Name  string
	Price float64
}

// create slice for products
var products []Product

func main() {
	var n, i int
	fmt.Println("How many product do you want to enter?")
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for i = 0; i < n; i++ {
		var id int
		var price float64
		fmt.Printf("\nProduct %d\n", i+1)

		fmt.Print("Enter ID: ")
		fmt.Scan(&id)

		scanner.Scan()

		fmt.Print("Enter Name: ")
		scanner.Scan()
		name := scanner.Text()

		fmt.Print("Enter Price: ")
		fmt.Scan(&price)

		product := Product{
			ID:    id,
			Name:  name,
			Price: price,
		}
		products = append(products, product)
	}
	output := ""
	for _, p := range products {
		output += fmt.Sprintf("%d,%s,%.2f\n", p.ID, p.Name, p.Price)
	}
	err := os.WriteFile("book-list.txt", []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing file", err)

		return
	}

	data, err := os.ReadFile("book-list.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("\n book List:")
	fmt.Println(string(data))
}
