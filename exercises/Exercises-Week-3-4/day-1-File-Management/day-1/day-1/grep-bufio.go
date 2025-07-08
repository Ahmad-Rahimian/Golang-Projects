// Summary: Searches for a keyword in each line of a file and prints matching lines.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("notes.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	var keyword string
	fmt.Print("Enter word to search: ")
	fmt.Scan(&keyword)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			fmt.Println("Found:", line)
		}
	}
}
