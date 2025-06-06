package main

import "fmt"

func main() {
	var n, m, temp int
	fmt.Println("enter 2 numbers :")
	fmt.Scan(&n)
	fmt.Scan(&m)
	max := 0
	if n > m {
		temp = n
		n = m
		m = temp
	}
	for i := 1; i <= m; i++ {
		if (n%i) == 0 && (m%i) == 0 {
			max = i
		}
	}
	fmt.Println(max)
}
