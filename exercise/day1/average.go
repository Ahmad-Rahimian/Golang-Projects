package main

import "fmt"

func main() {
	var n, sum float64
	var z float64
	fmt.Println("Please enter your number of numbers. :")
	fmt.Scan(&z)
	for i := 0; i < int(z); i++ {
		fmt.Println("please enter number-", i+1, ":")
		fmt.Scan(&n)
		sum += n
	}
	avg := (sum / z)
	fmt.Printf("numbers average : %v\n ", avg)
}
