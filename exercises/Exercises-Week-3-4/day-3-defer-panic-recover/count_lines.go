// Summary: Reads a file and counts its lines; uses panic if the file can't be read.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

func handle(filename string) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	return countLines(filename)
}

func main() {
	var filename string

	fmt.Println("enter file name :")
	fmt.Scan(&filename)

	count := handle(filename)

	fmt.Println("Line count : ", count)
}
