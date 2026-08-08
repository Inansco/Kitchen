package main

import (
	"html/template"
	"log"
	"net/http"
)

func render(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles("templates/" + page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "home.html")
}

func menu(w http.ResponseWriter, r *http.Request) {
	render(w, "menu.html")
}

func about(w http.ResponseWriter, r *http.Request) {
	render(w, "about.html")
}

func contact(w http.ResponseWriter, r *http.Request) {
	render(w, "contact.html")
}

func reservation(w http.ResponseWriter, r *http.Request) {
	render(w, "reservation.html")
}

func recommend(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		render(w, "recommend.html")
		return
	}

	food := r.FormValue("food")

	rec, err := GetRecommendation(food)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	t := template.Must(template.ParseFiles("templates/recommend.html"))

	t.Execute(w, rec)
}

func main() {

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", home)
	http.HandleFunc("/menu", menu)
	http.HandleFunc("/order", order)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/recommend", recommend)
	http.HandleFunc("/reservation", reservation)

	log.Println("Server started at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
