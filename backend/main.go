package main

import (
	"net/http"

	"MrOverflow.github.io/mortgage-underwriting/backend/db"
	"MrOverflow.github.io/mortgage-underwriting/backend/middlewares"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	defer db.Close()
	err := db.Initialize()

	if err != nil {
		panic("Failed to initialize database: " + err.Error())
	}

	r := mux.NewRouter()
	r.Use(middlewares.LoggerMiddleware)

	addRoutes(r)

	http.ListenAndServe(":8080", r)
}
