package main

import (
	"HarlanCinema/pkg/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

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
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	var page Page

	movies, err := app.services.MovieService.GetAllMovies()
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}

	page.User = getUserFromContext(r)
	page.Movies = movies
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
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	err = ts.Execute(w, page)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
}

func (app *application) allSeances(w http.ResponseWriter, r *http.Request) {
	var page Page

	seances, err := app.services.SeanceService.GetAllSeances()
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	page.User = getUserFromContext(r)
	page.Seances = seances

	files := []string{
		"./ui/html/all_seances.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	err = ts.Execute(w, page)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
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
	session, _ := app.sessions.Get(r, SessionName)
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
	email := r.Form.Get("email")
	fmt.Println(email)
	user := models.User{
		Username: username,
		Password: password,
		Email:    email,
	}
	user, err = app.services.UserService.RegisterUser(user)
	if err != nil {
		app.errorLog.Println("Some error occured in register POST", err)
		http.Error(w, "Internal server error", 500)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	// Fetch the session
	session, err := app.sessions.Get(r, SessionName)
	if err != nil {
		app.errorLog.Printf("Error fetching session: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if _, ok := session.Values["user_id"]; ok {
		delete(session.Values, "user_id")
		err = session.Save(r, w)
		if err != nil {
			app.errorLog.Printf("Error saving session: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
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

	user := getUserFromContext(r)

	tickets, err := app.services.TicketService.Repo.TicketRepository.FindAllTicketsByUserID(user.ID)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	for i := range tickets {
		mov, _ := app.services.MovieService.GetMovieById(tickets[i].Seance.MovieID)
		tickets[i].Seance.Movie = mov
	}
	page.Tickets = tickets
	page.User = user
	files := []string{
		"./ui/html/my_tickets.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	err = ts.Execute(w, page)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
}

func (app *application) aboutFilm(w http.ResponseWriter, r *http.Request) {
	var page Page

	movieId, _ := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	movie, err := app.services.MovieService.GetMovieById(movieId)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	page.User = getUserFromContext(r)
	page.Movie = movie
	seances, err := app.services.SeanceService.Repo.SeanceRepository.FindByMovieID(movieId)
	if err != nil {
		app.errorLog.Println("Error getting film", err)
		http.Error(w, "Internal Server error", 500)
		return
	}
	page.Seances = seances
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

func (app *application) getTicket(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.errorLog.Printf("Error parsing form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user := getUserFromContext(r)

	seance := r.FormValue("seanceID")
	fmt.Println(seance)
	seanceID, err := strconv.Atoi(seance)

	err = app.services.TicketService.GiveTicketForUser(user.ID, int64(seanceID))
	if err != nil {
		app.errorLog.Printf("Error parsing form: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/my-tickets", 303)
}
