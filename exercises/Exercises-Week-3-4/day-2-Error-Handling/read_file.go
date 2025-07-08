// Summary: Tries to read a file; if it doesn't exist, prints a friendly error message.
package main

import (
	"bufio"
	"fmt"
	"os"
)

var filename string

func main() {
	fmt.Println("please enter file name : ")
	fmt.Scan(&filename)
	file, err := os.Open(filename)
	if os.IsNotExist(err) {
		fmt.Println("file not found: ", filename)
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	file.Close()
}
