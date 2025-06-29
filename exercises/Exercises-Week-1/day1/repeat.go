package main

import "fmt"

func main() {
	var n int
	counter := make(map[int]int)
	for i := 0; i < 5; i++ {
		fmt.Println("please insert numbers ", i+1, ":")
		fmt.Scan(&n)
		counter[n]++
	}

	for key, value := range counter {
		fmt.Printf("number %v repeat %v \n", key, value)
	}
}
