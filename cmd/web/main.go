package main

import (
	"HarlanCinema/pkg/models"
	repo "HarlanCinema/pkg/repos"
	"HarlanCinema/pkg/services"
	"flag"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

const SessionName = "GSESS"

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	sessions *sessions.CookieStore
	services *services.Service
}

func main() {
	addr := flag.String("addr", ":3000", "HTTP network address")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	secret := os.Getenv("secret")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	sessionKey := secret
	store := sessions.NewCookieStore([]byte(sessionKey))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Movie{}, &models.Review{}, &models.Seance{}, &models.Ticket{})
	if err != nil {
		log.Fatal("Problems with auto migrating db", err)
	}

	repos := repo.NewRepository(db)

	srs := services.NewService(repos)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		sessions: store,
		services: srs,
	}

	r := app.routes()

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  r,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
