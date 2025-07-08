package main

import (
	"fmt"
)

type Student struct {
	Name    string
	ID      string
	Courses map[string]float64
}

var students = make(map[string]Student)

func addStudent() {
	var count int
	fmt.Print("How many students do you want to add? ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		var name, id string
		var courseCount int

		fmt.Printf("\n📘 Student %d\n", i+1)
		fmt.Print("Name: ")
		fmt.Scan(&name)

		fmt.Print("ID: ")
		fmt.Scan(&id)

		fmt.Print("How many courses? ")
		fmt.Scan(&courseCount)

		courses := make(map[string]float64)

		for j := 0; j < courseCount; j++ {
			var courseName string
			var score float64

			fmt.Printf("  Course %d name: ", j+1)
			fmt.Scan(&courseName)

			fmt.Print("  Score: ")
			fmt.Scan(&score)

			courses[courseName] = score
		}

		students[id] = Student{
			Name:    name,
			ID:      id,
			Courses: courses,
		}
	}
}

func showStudent() {
	var id string
	fmt.Print("Enter student ID to search: ")
	fmt.Scan(&id)

	if student, found := students[id]; found {
		fmt.Printf("Name: %s\n", student.Name)
		fmt.Println("Courses and Scores:")
		var sum float64
		for name, score := range student.Courses {
			fmt.Printf("  %s: %.2f\n", name, score)
			sum += score
		}
		average := sum / float64(len(student.Courses))
		fmt.Printf("Average Score: %.2f\n", average)
	} else {
		fmt.Println("Student not found.")
	}
}

func main() {
	for {
		fmt.Println("\n====== MENU ======")
		fmt.Println("1. Add student")
		fmt.Println("2. Show student and average")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addStudent()
		case 2:
			showStudent()
		case 3:
			fmt.Println("Exiting program.")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
