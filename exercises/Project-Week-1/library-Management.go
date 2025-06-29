// Summary: A simple modular library system with book and member management using structs, maps, and interfaces.

package main

import "fmt"

// structs
type Book struct {
	ID     string
	Title  string
	Author string
}
type Member struct {
	ID    string
	Name  string
	Email string
}
type Library struct {
	Books   map[string]Book
	Members map[string]Member
}

type BookManager interface {
	AddBook(book Book)
	RemoveBook(id string)
	SearchBook(id string)
}

type MemberManager interface {
	AddMember(member Member)
	RemoveMember(id string)
	SearchMember(id string)
}

func (l *Library) AddBook(b Book) {
	l.Books[b.ID] = b
	fmt.Println("book added")
}

func (l *Library) AddMember(m Member) {
	l.Members[m.ID] = m
	fmt.Println("Member added")
}

func (l *Library) RemoveBook(id string) {
	delete(l.Books, id) // delete(map,key)
	fmt.Println("book removed")
}

func (l *Library) RemoveMember(id string) {
	delete(l.Members, id) // delete(map,key)
	fmt.Println("member removed")
}

func (l *Library) SearchBook(id string) {
	if book, found := l.Books[id]; found {
		fmt.Printf("Book found - Title : %s , Author : %s \n", book.Title, book.Author)
	} else {
		fmt.Println("Book not found")
	}
}

func (l *Library) SearchMember(id string) {
	if member, found := l.Members[id]; found {
		fmt.Printf("Member found - Name : %s , Email: : %s \n", member.Name, member.Email)
	} else {
		fmt.Println("Member not found")
	}
}

func main() {
	library := Library{
		Books:   make(map[string]Book),
		Members: make(map[string]Member),
	}

	for {
		fmt.Println("\n--- Library Menu ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Search Book")
		fmt.Println("3. Remove Book")
		fmt.Println("4. Add Member")
		fmt.Println("5. Search Member")
		fmt.Println("6. Remove Member")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, title, author string
			fmt.Println("Book ID :")
			fmt.Scan(&id)
			fmt.Println("Book Title :")
			fmt.Scan(&title)
			fmt.Println("Book Author :")
			fmt.Scan(&author)
			book := Book{ID: id, Title: title, Author: author}
			library.AddBook(book)

		case 2:
			var id string
			fmt.Println("Book ID :")
			fmt.Scan(&id)
			library.SearchBook(id)

		case 3:
			var id string
			fmt.Println("Book ID :")
			fmt.Scan(&id)
			library.RemoveBook(id)

		case 4:
			var name, id, email string
			fmt.Println("Member ID :")
			fmt.Scan(&id)
			fmt.Println("Member name :")
			fmt.Scan(&name)
			fmt.Println("Member Email :")
			fmt.Scan(&email)
			member := Member{ID: id, Name: name, Email: email}
			library.AddMember(member)

		case 5:
			var id string
			fmt.Println("Member ID :")
			fmt.Scan(&id)
			library.SearchMember(id)

		case 6:
			var id string
			fmt.Println("Member ID :")
			fmt.Scan(&id)
			library.RemoveMember(id)

		case 7:
			fmt.Println("Goodbye")
			return
		}
	}
}
