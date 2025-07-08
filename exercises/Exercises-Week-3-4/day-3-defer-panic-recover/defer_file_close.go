// Summary: Opens a file and ensures it is closed using defer, regardless of errors.

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Fprintln(file, "your log message")
	defer file.Close()
}
