package main

import "fmt"

func countEvenOdd() (int, int) {
	var number, count int
	var countEven, countOdd int
	fmt.Println("Please enter how many numbers : ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		fmt.Printf("enter number %d: ", i+1)
		fmt.Scan(&number)

		if number%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}
	return countEven, countOdd
}

func main() {
	even, odd := countEvenOdd()
	fmt.Printf("even count is: %d\n", even)
	fmt.Printf("odd count is: %d\n", odd)
}
