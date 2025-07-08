// Summary: Simulates a simple logic error system that logs issues using panic and recover.
package main

import (
	"fmt"
)

func processData(data string) {
	if data == "" {
		panic("data cannot be empty")
	}
	if data == "bad" {
		panic("invalid data value")
	}
}

func handle(data string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	processData(data)
}

func main() {
	var data string
	fmt.Println("enter phasrase: ")
	fmt.Scan(&data)
	handle(data)
}
