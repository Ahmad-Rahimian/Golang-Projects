// Summary: A simple digital clock that prints current time every second using a Goroutine.
package main

import (
	"fmt"
	"time"
)

func showTime() {
	for i := 0; i <= 10; i++ {
		fmt.Println(time.Now().Format("15:04:05"))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go showTime()
	time.Sleep(10 * time.Second)
}
