package main

import (
	"fmt"
)

func main() {
	var sliceBig []int
	var sliceSmall []int

	var n int
	for i := 0; i < 10; i++ {
		fmt.Printf("please enter number %d :\n", i+1)
		fmt.Scan(&n)
		sliceBig = append(sliceBig, n)
	}

	var m int
	for j := 0; j < 5; j++ {
		fmt.Printf("please enter number %d:\n", j+1)
		fmt.Scan(&m)
		sliceSmall = append(sliceSmall, m)
	}

	found := false
	for i := 0; i <= len(sliceBig)-len(sliceSmall); i++ {
		match := true
		for j := 0; j < len(sliceSmall); j++ {
			if sliceBig[i+j] != sliceSmall[j] {
				match = false
				break
			}
		}
		if match {
			found = true
			break
		}
	}

	if found {
		fmt.Println("The second array was found in the first array.")
	} else {
		fmt.Println("The second array was NOT found in the first array.")
	}
}
