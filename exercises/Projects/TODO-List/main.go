// Summary: Full-featured Todo app with create, update, and delete functionality.

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// ساختار هر Todo
type Todo struct {
	Title   string
	Content string
}

// حافظه موقت برای نگهداری Todoها
var todos []Todo

// نمایش لیست Todoها
func showTodos(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, todos)
}

// نمایش فرم ایجاد Todo
func showCreateForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("create.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

// ایجاد Todo جدید
func createTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/create", http.StatusSeeOther)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")

	todos = append(todos, Todo{Title: title, Content: content})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// حذف یک Todo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	indexStr := r.URL.Query().Get("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(todos) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}
	todos = append(todos[:index], todos[index+1:]...)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// نمایش فرم ویرایش و پردازش آن
func editTodoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		indexStr := r.URL.Query().Get("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(todos) {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		tmpl, err := template.ParseFiles("edit.html")
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, struct {
			Index int
			Todo  Todo
		}{
			Index: index,
			Todo:  todos[index],
		})

	} else if r.Method == http.MethodPost {
		indexStr := r.FormValue("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(todos) {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		todos[index].Title = r.FormValue("title")
		todos[index].Content = r.FormValue("content")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	indexStr := r.FormValue("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(todos) {
		http.NotFound(w, r)
		return
	}

	// گرفتن داده جدید از فرم
	title := r.FormValue("title")
	content := r.FormValue("content")

	todos[index].Title = title
	todos[index].Content = content

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// راه‌اندازی سرور و مسیرها
func main() {
	http.HandleFunc("/", showTodos)
	http.HandleFunc("/create", showCreateForm)
	http.HandleFunc("/add", createTodo)
	http.HandleFunc("/delete", deleteTodo)
	http.HandleFunc("/edit", editTodoHandler)
	http.HandleFunc("/update", updateTodo)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
