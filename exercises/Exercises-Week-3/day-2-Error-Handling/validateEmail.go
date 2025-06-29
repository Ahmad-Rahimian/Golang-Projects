package main

import (
	"errors"
	"fmt"
	"strings"
)

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("invalid email: missing '@'")
	}
	return nil
}

func main() {
	var email string
	fmt.Println("please enter your email : ")
	fmt.Scan(&email)

	err := validateEmail(email)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("Email registered successfully")
}
