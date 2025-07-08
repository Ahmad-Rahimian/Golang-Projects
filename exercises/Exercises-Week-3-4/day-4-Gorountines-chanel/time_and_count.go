// Summary: Runs two goroutines: one prints current time every second, the other counts numbers.
package main

import (
	"fmt"
	"time"
)

func showTime() {
	for i := 0; i <= 10; i++ {
		fmt.Println("time", time.Now().Format("15:04:05"))
		time.Sleep(1 * time.Second)
	}
}

func countNumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println("count", i)
		time.Sleep(800 * time.Millisecond)
	}
}

func main() {
	go showTime()
	go countNumbers()
	time.Sleep(6 * time.Second)
}
