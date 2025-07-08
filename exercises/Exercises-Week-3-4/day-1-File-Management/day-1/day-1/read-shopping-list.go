package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, i int
	fmt.Println("How many items do you want to enter?")
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var items string
	for i = 0; i < n; i++ {
		fmt.Println("please enter title for items?")
		scanner.Scan()
		item := scanner.Text()
		items += item + "\n"
	}
	err := os.WriteFile("shopping-list.txt", []byte(items), 0644)
	if err != nil {
		fmt.Println("Error writing file", err)

		return
	}

	data, err := os.ReadFile("shopping-list.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("\n Shopping List:")
	fmt.Println(string(data))
}
