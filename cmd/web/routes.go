package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() *mux.Router {
	routes := mux.NewRouter()
	file := routes.PathPrefix("/").Subrouter()
	r := routes.PathPrefix("/").Subrouter()
	fs := http.FileServer(http.Dir("./ui/static"))
	file.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	file.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./uploads"))))

	r.Use(app.logRequest)
	r.Use(app.userAuth)
	r.Use(app.RecoverPanic)
	r.Use(app.SecureHeaders)

	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/all-seances", app.allSeances).Methods("GET")
	r.HandleFunc("/about-film", app.aboutFilm).Methods("GET")
	r.HandleFunc("/login", app.loginPage).Methods("GET")
	r.HandleFunc("/register", app.registerPage).Methods("GET")
	r.HandleFunc("/registration", app.register).Methods("POST")
	r.HandleFunc("/loginform", app.login).Methods("POST")

	authRoutes := r.PathPrefix("/").Subrouter()
	authRoutes.Use(app.RequireAuthentication)
	authRoutes.HandleFunc("/my-tickets", app.myTickets).Methods("GET")
	authRoutes.HandleFunc("/logout", app.logout).Methods("GET")
	authRoutes.HandleFunc("/getTicket", app.getTicket).Methods("POST")

	adminRoutes := authRoutes.PathPrefix("/").Subrouter()
	adminRoutes.Use(app.RequireAdmin)
	adminRoutes.HandleFunc("/admin", app.adminPage).Methods("GET")
	adminRoutes.HandleFunc("/admin/movie", app.createMovie).Methods("POST")
	adminRoutes.HandleFunc("/admin/seance", app.createSeance).Methods("POST")

	return routes
}
