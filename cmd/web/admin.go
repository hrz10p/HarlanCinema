package main

import (
	"HarlanCinema/pkg/models"
	"net/http"
	"strconv"
	"time"
)

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(20 << 20)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}

	err = r.ParseForm()
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}

	title := r.FormValue("title")
	desc := r.FormValue("description")
	rating := r.FormValue("rating")

	raing, err := strconv.Atoi(rating)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	_, header, err := r.FormFile("image")
	if header != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}

	fileURL, err := app.services.FileService.SaveFile(header)

	movie := models.Movie{
		Title:       title,
		Description: desc,
		Rating:      float64(raing),
		ImageUrl:    fileURL,
	}

	movie, err = app.services.MovieService.CreateMovie(movie)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}

	http.Redirect(w, r, "/admin", 303)

}

func (app *application) createSeance(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}

	location := r.FormValue("location")
	date := r.FormValue("date")
	movie := r.FormValue("movie")

	movint, err := strconv.Atoi(movie)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	times, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	seance := models.Seance{
		MovieID:  int64(movint),
		Date:     times,
		Location: location,
		Movie:    models.Movie{},
	}

	_, err = app.services.SeanceService.CreateSeance(seance)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}

	http.Redirect(w, r, "/admin", 303)
}
