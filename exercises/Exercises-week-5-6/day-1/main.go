package main

import (
	"html/template"
	"net/http"
	"sync"
)

// Task ساختار مربوط به تسک
type Task struct {
	ID      int
	Title   string
	Details string
}

// tasks لیست تسک‌ها (در حافظه ذخیره میشه)
var tasks []Task

// mutex برای جلوگیری از race condition هنگام دسترسی به تسک‌ها
var mutex sync.Mutex

// tmpl قالب HTML صفحه
var tmpl = template.Must(template.ParseFiles("template.html"))

func main() {
	http.HandleFunc("/", taskHandler)
	http.ListenAndServe(":8080", nil)
}

// taskHandler هندلر اصلی برای نمایش و افزودن تسک‌ها
func taskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		details := r.FormValue("details")

		mutex.Lock()
		newTask := Task{
			ID:      len(tasks) + 1,
			Title:   title,
			Details: details,
		}
		tasks = append(tasks, newTask)
		mutex.Unlock()
	}

	// نمایش صفحه با تمام تسک‌ها
	tmpl.Execute(w, tasks)
}
