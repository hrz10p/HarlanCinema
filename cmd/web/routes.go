package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/all-seances", app.allSeances)
	mux.HandleFunc("/login", app.loginPage)
	mux.HandleFunc("/register", app.registerPage)
	mux.HandleFunc("/my-tickets", app.myTickets)
	mux.HandleFunc("/about-film", app.aboutFilm)

	fs := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	return mux
}
