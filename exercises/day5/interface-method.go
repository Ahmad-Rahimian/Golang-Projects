// create interface and method
package main

import "fmt"

type Animal interface {
	speek() string
}

type (
	Dog  struct{}
	Cat  struct{}
	Bird struct{}
)

func (d Dog) speek() string {
	return "hop hop"
}

func (c Cat) speek() string {
	return "miu miu"
}

func (b Bird) speek() string {
	return "jik jik"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Bird{}}
	for _, a := range animals {
		fmt.Println(a.speek())
	}
}
