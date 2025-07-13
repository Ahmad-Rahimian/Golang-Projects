package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	_ "github.com/lib/pq"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

var contacts []Contact

func showContact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("contacts.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, contacts)
}

func showCreateFrom(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("create.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func createContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/create", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")

	contacts = append(contacts, Contact{Name: name, Phone: phone, Email: email})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	indexStr := r.URL.Query().Get("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(contacts) {
		http.Error(w, "invalid index", http.StatusSeeOther)
		return
	}
	contacts = append(contacts[:index], contacts[index+1:]...)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		indexStr := r.URL.Query().Get("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(contacts) {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}
		tmpl, err := template.ParseFiles("edit.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, struct {
			Index   int
			Contact Contact
		}{
			Index:   index,
			Contact: contacts[index],
		})
	} else if r.Method == http.MethodPost {
		indexStr := r.FormValue("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(contacts) {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		contacts[index].Name = r.FormValue("name")
		contacts[index].Phone = r.FormValue("phone")
		contacts[index].Email = r.FormValue("email")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	indexStr := r.FormValue("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(contacts) {
		http.NotFound(w, r)
		return
	}

	// گرفتن داده جدید از فرم
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	email := r.FormValue("email")

	contacts[index].Name = name
	contacts[index].Phone = phone
	contacts[index].Email = email

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", showContact)
	http.HandleFunc("/create", showCreateFrom)
	http.HandleFunc("/add", createContact)
	http.HandleFunc("/delete", deleteContact)
	http.HandleFunc("/edit", editContactHandler)
	http.HandleFunc("/update", updateContact)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
