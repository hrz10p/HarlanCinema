package main

import (
	"HarlanCinema/pkg/models"
	"HarlanCinema/pkg/utils/logger"
	"context"
	"net/http"
)

type contextKey struct {
	name string
}

var userCtxKey = contextKey{name: "user"}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) userAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := app.sessions.Get(r, SessionName)
		if err != nil {
			app.errorLog.Println("Error fetching session:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			next.ServeHTTP(w, r)
		}

		user, err := app.services.UserService.GetUserByID(id.(string))
		if err != nil {
			app.errorLog.Println("Error fetching user:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}

func (app *application) RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromContext(r)
		if (user == models.User{}) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				logger.GetLogger().Error(err.(error).Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func getUserFromContext(r *http.Request) models.User {
	user, ok := r.Context().Value(userCtxKey).(models.User)
	if !ok {
		logger.GetLogger().Info("User is not authenticated")
		return models.User{}
	}
	return user
}
