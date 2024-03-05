package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() *mux.Router {
	r := mux.NewRouter()

	r.Use(app.logRequest)
	r.Use(app.userAuth)
	r.Use(app.RecoverPanic)
	r.Use(app.SecureHeaders)

	authRoutes := r.PathPrefix("/").Subrouter()
	authRoutes.Use(app.RequireAuthentication)
	authRoutes.HandleFunc("/my-tickets", app.myTickets)

	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/all-seances", app.allSeances).Methods("GET")
	r.HandleFunc("/about-film", app.aboutFilm).Methods("GET")
	r.HandleFunc("/login", app.loginPage).Methods("GET")
	r.HandleFunc("/register", app.registerPage).Methods("GET")

	// Static file server
	fs := http.FileServer(http.Dir("./ui/static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./uploads"))))

	return r
}
