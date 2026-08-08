package main

import (
	"html/template"
	"net/http"
)

type Order struct {
	Name     string
	Phone    string
	Meal     string
	Quantity string
}

func order(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		render(w, "order.html")
		return
	}

	o := Order{
		Name:     r.FormValue("name"),
		Phone:    r.FormValue("phone"),
		Meal:     r.FormValue("meal"),
		Quantity: r.FormValue("quantity"),
	}

	t := template.Must(
		template.ParseFiles("templates/success.html"),
	)

	t.Execute(w, o)
}
