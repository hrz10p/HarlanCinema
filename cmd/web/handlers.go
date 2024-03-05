package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Movie struct {
	ImageUrl    string
	Title       string
	Rating      float64
	ID          int
	Description string
}
type Cinema struct {
	ImageUrl string
	Title    string
	Rating   float64
}
type Ticket struct {
	ImageUrl   string
	Title      string
	Date       string
	Time       string
	Cinema     string
	TicketType string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	premiers := []Movie{
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
	}
	rentMovies := []Movie{
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
	}
	cinemas := []Cinema{
		{ImageUrl: "https://avatars.mds.yandex.net/get-altay/6446898/2a000001841f6855cc48a43506d62c061537/orig", Title: "Cool cinema", Rating: 5.0},
	}
	type TMPL struct {
		Premiers   []Movie
		RentMovies []Movie
		Cinemas    []Cinema
	}
	tmpl := TMPL{Cinemas: cinemas, Premiers: premiers, RentMovies: rentMovies}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//app.serverError(w, err)
		return
	}
	err = ts.Execute(w, tmpl)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) allSeances(w http.ResponseWriter, r *http.Request) {
	premiers := []Movie{
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
		{ImageUrl: "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg", Title: "Interstellar", Rating: 5.0, ID: 1},
	}
	files := []string{
		"./ui/html/all_seances.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//app.serverError(w, err)
		return
	}
	err = ts.Execute(w, premiers)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) loginPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/login.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) registerPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/register.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//app.serverError(w, err)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) myTickets(w http.ResponseWriter, r *http.Request) {
	tickets := []Ticket{
		{
			ImageUrl:   "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg",
			Title:      "Interstellar",
			Date:       "01.01.2024",
			Time:       "00:00",
			Cinema:     "Cool cinema",
			TicketType: "Child",
		},
		{
			ImageUrl:   "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg",
			Title:      "Interstellar",
			Date:       "01.01.2024",
			Time:       "00:00",
			Cinema:     "Cool cinema",
			TicketType: "Child",
		},
		{
			ImageUrl:   "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg",
			Title:      "Interstellar",
			Date:       "01.01.2024",
			Time:       "00:00",
			Cinema:     "Cool cinema",
			TicketType: "Child",
		},
		{
			ImageUrl:   "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg",
			Title:      "Interstellar",
			Date:       "01.01.2024",
			Time:       "00:00",
			Cinema:     "Cool cinema",
			TicketType: "Child",
		},
	}
	files := []string{
		"./ui/html/my_tickets.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//app.serverError(w, err)
		return
	}
	err = ts.Execute(w, tickets)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) aboutFilm(w http.ResponseWriter, r *http.Request) {
	tickets := Movie{
		ImageUrl:    "https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_.jpg",
		Title:       "Interstellar",
		Rating:      5.0,
		Description: "Cool film, i like so much",
	}
	files := []string{
		"./ui/html/about_film.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		//app.serverError(w, err)
		return
	}
	err = ts.Execute(w, tickets)
	if err != nil {
		fmt.Println("Error")
		return
	}
}
