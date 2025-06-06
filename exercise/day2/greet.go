package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello ,%s\n", name)
}

func main() {
	var name string
	fmt.Println("what is your name ?")
	fmt.Scan(&name)
	greet(name)
}
