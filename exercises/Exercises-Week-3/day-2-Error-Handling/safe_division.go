package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	var a, b int
	fmt.Println("please enter number one")
	fmt.Scan(&a)
	fmt.Println("please enter number two")
	fmt.Scan(&b)

	result, err := division(a, b)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
	fmt.Printf("result: %v", result)
}
