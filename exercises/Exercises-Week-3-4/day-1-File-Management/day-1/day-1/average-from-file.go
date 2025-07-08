package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, i int
	fmt.Println("How many score do you want to enter?")
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	var scores string
	for i = 0; i < n; i++ {
		fmt.Println("please enter score", i+1, ":")
		scanner.Scan()
		score := scanner.Text()
		scores += score + "\n"
	}
	err := os.WriteFile("scores.txt", []byte(scores), 0644)
	if err != nil {
		fmt.Println("Error writing file", err)

		return
	}

	data, err := os.ReadFile("scores.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")
	var sum int
	var count int
	for _, line := range lines {
		if line == "" {
			continue
		}
		score, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("invalid number", line)
			continue
		}
		sum += score
		count++
	}
	if count > 0 {
		average := float64(sum) / float64(count)
		fmt.Println("average :", average)
	} else {
		fmt.Println("No valid scores to calculate average.")
	}
}
