package main

import (
	"HarlanCinema/pkg/models"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

func (app *application) adminPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/admin.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	var page Page
	movies, err := app.services.MovieService.GetAllMovies()
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	page.Movies = movies
	err = ts.Execute(w, page)
	if err != nil {
		app.errorLog.Println(err)
		http.Error(w, "Internal Server Error", 500)
	}
}

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
	imageUrl := r.FormValue("image_url")

	raing, err := strconv.Atoi(rating)
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}
	_, header, err := r.FormFile("image")
	if err != nil {
		app.errorLog.Println("Something wrong", err)
		http.Error(w, "Internal Server error", 500)
	}

	fileURL, err := app.services.FileService.SaveFile(header)

	movie := models.Movie{
		Title:       title,
		Description: desc,
		Rating:      float64(raing),
		ImageUrl:    imageUrl,
	}
	fmt.Println(fileURL)
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
	const layout = "2006-01-02T15:04"
	times, err := time.Parse(layout, date)
	if err != nil {
		app.errorLog.Printf("Error parsing date: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
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
