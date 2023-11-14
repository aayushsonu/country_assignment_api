package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "country_assignment_api/docs"
)

// @title API_Assignment | Swagger
// @version 1.0
// @description API Documentation for all the endpoints
// @termsOfService https://example.com/terms/
// @host localhost:8080
// @basePath /api/v1

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/auth", AuthHandler).Methods("POST")
	r.Handle("/api/v1/country", AuthMiddleware(http.HandlerFunc(CountryDetailsHandler))).Methods("GET")
	r.Handle("/api/v1/countries", AuthMiddleware(http.HandlerFunc(CountriesListHandler))).Methods("GET")
	r.Handle("/api/v1/countries/filter", AuthMiddleware(http.HandlerFunc(CountriesFilterListHandler))).Methods("GET")

	// Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The URL (or route) for the Swagger JSON endpoint
	))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}