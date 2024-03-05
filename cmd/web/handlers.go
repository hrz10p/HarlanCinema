package main

import (
	"HarlanCinema/pkg/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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

type Page struct {
	User    models.User
	Movies  []models.Movie
	Seances []models.Seance
	Tickets []models.Ticket
	Reviews []models.Review
	Movie   models.Movie
	Ticket  models.Ticket
	Seance  models.Seance
	Review  models.Review

	Premiers   []models.Seance
	RentMovies []models.Movie
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	var page Page

	movies, err := app.services.MovieService.GetAllMovies()
	seances, err := app.services.SeanceService.GetAllSeances()

	page.User = getUserFromContext(r)
	page.RentMovies = movies
	page.Premiers = seances

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
	err = ts.Execute(w, page)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) allSeances(w http.ResponseWriter, r *http.Request) {
	var page Page

	seances, err := app.services.SeanceService.GetAllSeances()

	page.User = getUserFromContext(r)
	page.Seances = seances

	files := []string{
		"./ui/html/all_seances.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		return
	}
	err = ts.Execute(w, page)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) loginPage(w http.ResponseWriter, r *http.Request) {
	session, _ := app.sessions.Get(r, SessionName)

	flashes := session.Flashes()
	session.Save(r, w)

	files := []string{
		"./ui/html/login.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	data := map[string]interface{}{
		"Flashes": flashes,
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Printf("Error parsing form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	session, _ := app.sessions.Get(r, "session-name")
	user, err := app.services.UserService.AuthenticateUser(username, password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			session.AddFlash("Invalid username or password")
			session.Save(r, w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else {
			app.errorLog.Println("Some error occured in login POST", err)
			http.Error(w, "Internal server error", 500)
			return
		}

	}

	session.Values["user_id"] = user.ID
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Printf("Error parsing form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	username := r.Form.Get("username")
	password := r.Form.Get("password")
	user := models.User{
		Username: username,
		Password: password,
	}
	user, err = app.services.UserService.RegisterUser(user)
	if err != nil {
		app.errorLog.Println("Some error occured in register POST", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) registerPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/register.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println("Unable to parse tmpls")
		http.Error(w, "Internal server error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) myTickets(w http.ResponseWriter, r *http.Request) {
	var page Page

	tickets, err := app.services.TicketService.GetAllTickets()

	page.User = getUserFromContext(r)
	page.Tickets = tickets

	files := []string{
		"./ui/html/my_tickets.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		return
	}
	err = ts.Execute(w, page)
	if err != nil {
		fmt.Println("Error")
		return
	}
}

func (app *application) aboutFilm(w http.ResponseWriter, r *http.Request) {
	var page Page

	movieId, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	movie, err := app.services.MovieService.GetMovieById(movieId)

	page.User = getUserFromContext(r)
	page.Movie = movie

	if err != nil {
		app.errorLog.Println("Error getting film", err)
		http.Error(w, "Internal Server error", 500)
		return
	}
	files := []string{
		"./ui/html/about_film.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		return
	}
	err = ts.Execute(w, page)
	if err != nil {
		return
	}
}
