package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type Note struct {
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

var notes []Note

func showNoteList(w http.ResponseWriter, r *http.Request) {
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].CreatedAt.After(notes[j].CreatedAt)
	})

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, notes)
}

func showCreateForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("create.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func addNote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/create", http.StatusSeeOther)
		return
	}

	now := time.Now()
	newNote := Note{
		Title:     r.FormValue("title"),
		Body:      r.FormValue("body"),
		CreatedAt: now,
		UpdatedAt: now,
	}
	notes = append(notes, newNote)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	indexStr := r.FormValue("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(notes) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}
	notes = append(notes[:index], notes[index+1:]...)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editNote(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		indexStr := r.URL.Query().Get("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(notes) {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		tmpl, err := template.ParseFiles("edit.html")
		if err != nil {
			http.Error(w, "Template Error", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, struct {
			Index int
			Note  Note
		}{
			Index: index,
			Note:  notes[index],
		})
	} else if r.Method == http.MethodPost {
		indexStr := r.FormValue("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(notes) {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		body := r.FormValue("body")

		notes[index].Title = title
		notes[index].Body = body
		notes[index].UpdatedAt = time.Now()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", showNoteList)
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			showCreateForm(w, r)
		} else if r.Method == http.MethodPost {
			addNote(w, r)
		}
	})
	http.HandleFunc("/delete", deleteNote)
	http.HandleFunc("/edit", editNote)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
