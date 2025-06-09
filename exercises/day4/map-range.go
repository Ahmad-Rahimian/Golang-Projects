package main

import "fmt"

func addStudent() {
	var i, number, score int
	var name string
	student := make(map[string]int)
	fmt.Println("please enter how many student?")
	fmt.Scan(&number)
	for i = 0; i < number; i++ {
		fmt.Printf("please enter student name %d:\n", i+1)
		fmt.Scan(&name)
		fmt.Printf("please enter student score %d:\n", i+1)
		fmt.Scan(&score)
		student[name] = score

	}
	fmt.Println("student list :")
	for name, score := range student {
		fmt.Printf("Name: %s | Score: %d\n", name, score)
	}
}

func main() {
	addStudent()
}
