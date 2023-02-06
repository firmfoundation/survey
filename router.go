package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	/*
	 For debugging only
	 comment it
	*/
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 1236})
	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}

func router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	//protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))

		// Handle valid / invalid tokens.
		r.Use(jwtauth.Authenticator)

	})

	r.Group(func(r chi.Router) {
		r.Method("GET", "/", Handler(Index))

		r.Method("POST", "/surveys", Handler(HandleCreateSurvey))

		//indicator
		r.Method("POST", "/indicators", Handler(HandleCreateIndicator))

		// question
		r.Method("POST", "/questions", Handler(HandleCreateQuestion))
		r.Method("GET", "/questions", Handler(HandleGetSurveyQuestion))

		//survey journals
		r.Method("POST", "/survey/journals", Handler(HandleCreateSurveyJournal))
		r.Method("POST", "/survey/results", Handler(HandleSurveyResult))

		//user
		r.Method("POST", "/users", Handler(HandleCreateUser))

		//execute
		r.Method("GET", "/exe", Handler(HandleExe))

	})

	//public folder for assets
	fs := http.FileServer(http.Dir(servicePath() + "/static/img"))
	r.Handle("/static/img/*", http.StripPrefix("/static/img/", fs))

	return r
}

func servicePath() string {
	ex, err := os.Executable()

	if err != nil {
		log.Fatal(err)
	}
	return filepath.Dir(ex)
}
