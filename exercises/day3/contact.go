package main

import (
	"fmt"
)

type Contact struct {
	Name        string
	PhoneNumber int
}

// Add new person
func addPerson(name string, phonenumber int) Contact {
	return Contact{Name: name, PhoneNumber: phonenumber}
}

// Update exact phonenumber
func updatePhone(p *Contact, newPhonenumber int) {
	p.PhoneNumber = newPhonenumber
}

// Update contact name
func updateName(p *Contact, newName string) {
	p.Name = newName
}

// Find person by name
func findPersonByName(contacts []Contact, name string) *Contact {
	for i := range contacts {
		if contacts[i].Name == name {
			return &contacts[i]
		}
	}
	return nil
}

// Print all contacts
func printcontacts(contacts []Contact) {
	if len(contacts) == 0 {
		fmt.Println("📭 No person found.")
		return
	}
	fmt.Println("\n Contact List:")
	for i, p := range contacts {
		fmt.Printf("%d. Name: %s | PhoneNumber: %d\n", i+1, p.Name, p.PhoneNumber)
	}
}

func main() {
	var contacts []Contact
	for {
		fmt.Println("\n====== MENU ======")
		fmt.Println("1. Add person")
		fmt.Println("2. Show contacts")
		fmt.Println("3. Update phonenumber")
		fmt.Println("4. Update contactname")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var phonenumber int
			var number, i int
			fmt.Println("Please enter how many numbers :")
			fmt.Scan(&number)
			for i = 0; i < number; i++ {
				fmt.Print("Enter name: ")
				fmt.Scan(&name)
				fmt.Print("Enter phonenumber: ")
				fmt.Scan(&phonenumber)
				contacts = append(contacts, addPerson(name, phonenumber))
				fmt.Println("✅ Contact added!")
			}
		case 2:
			printcontacts(contacts)

		case 3:
			var name string
			fmt.Print("Enter name to update phonenumber: ")
			fmt.Scan(&name)
			p := findPersonByName(contacts, name)
			if p != nil {
				var newPhonenumber int
				fmt.Print("Enter new phonenumber: ")
				fmt.Scan(&newPhonenumber)
				updatePhone(p, newPhonenumber)
				fmt.Println("✅ PhoneNumber updated.")
			} else {
				fmt.Println("❌ Contact not found.")
			}

		case 4:
			var name string
			fmt.Print("Enter name to update contact name: ")
			fmt.Scan(&name)
			p := findPersonByName(contacts, name)
			if p != nil {
				var newName string
				fmt.Print("Enter new name: ")
				fmt.Scan(&newName)
				updateName(p, newName)
				fmt.Println("✅ contact name updated.")
			} else {
				fmt.Println("❌ Contact not found.")
			}

		case 5:
			fmt.Println("👋 Exiting program.")
			return

		default:
			fmt.Println("❌ Invalid choice. Try again.")
		}
	}
}
