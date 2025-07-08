package main

import (
	"errors"
	"fmt"
)

func submitScore(score float64) error {
	if score < 0 {
		return errors.New("score cannot be negative")
	}
	fmt.Println("score registerd successfully")
	return nil
}

func main() {
	var score float64
	fmt.Println("please enter score : ")
	fmt.Scan(&score)
	err := submitScore(score)
	if err != nil {
		fmt.Println("error : ", err)
		return
	}
}
