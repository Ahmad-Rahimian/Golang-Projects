// Summary: A simple login system that checks username and password with custom error handling.

package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	storedUsername = "admin"
	storedPassword = "1234"
)

func Login(username, password string) error {
	if strings.ToLower(username) != storedUsername {
		return errors.New("invalid username")
	}
	if password != storedPassword {
		return errors.New("invalid Password")
	}
	return nil
}

func main() {
	var username, password string
	fmt.Println("Enter username :")
	fmt.Scan(&username)
	fmt.Println("Enter password :")
	fmt.Scan(&password)
	err := Login(username, password)
	if err != nil {
		fmt.Println("Login failed", err.Error())
		return
	}
	fmt.Println("Login successful")
}
