package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Product struct {
	Name        string
	Price       float64
	Description string
}

var products []Product

func showProductList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, products)
}

func showCreateForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("create.html")
	if err != nil {
		http.Error(w, "Template Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/create", http.StatusSeeOther)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		http.Error(w, "Invalid Price", http.StatusBadRequest)
		return
	}

	newProduct := Product{Name: name, Description: description, Price: price}
	products = append(products, newProduct)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	indexStr := r.FormValue("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil || index < 0 || index >= len(products) {
		http.Error(w, "Invalid index", http.StatusBadRequest)
		return
	}
	products = append(products[:index], products[index+1:]...)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		indexStr := r.URL.Query().Get("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(products) {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		tmpl, err := template.ParseFiles("edit.html")
		if err != nil {
			http.Error(w, "Template Error", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, struct {
			Index   int
			Product Product
		}{
			Index:   index,
			Product: products[index],
		})
	} else if r.Method == http.MethodPost {
		indexStr := r.FormValue("index")
		index, err := strconv.Atoi(indexStr)
		if err != nil || index < 0 || index >= len(products) {
			http.Error(w, "Invalid index", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		description := r.FormValue("description")
		priceStr := r.FormValue("price")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			http.Error(w, "Invalid Price", http.StatusBadRequest)
			return
		}

		products[index].Name = name
		products[index].Description = description
		products[index].Price = price

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", showProductList)
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			showCreateForm(w, r)
		} else if r.Method == http.MethodPost {
			addProduct(w, r)
		}
	})
	http.HandleFunc("/delete", deleteProduct)
	http.HandleFunc("/edit", editProduct)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
