package main

import (
	"fmt"
	"github.com/MasahiroYoshiichi/auth/cognito/cognitotoken"
	"github.com/MasahiroYoshiichi/auth/cognito/handlers"
	"github.com/rs/cors"
	"log"

	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
	router.HandleFunc("/confirm-signup", handlers.ConfirmSignUpHandler).Methods("POST")
	router.HandleFunc("/signin", handlers.SignInHandler).Methods("POST")
	router.Handle("/signout", cognitotoken.Middleware(http.HandlerFunc(handlers.SignOutHandler))).Methods("POST")

	// Configure CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Apply CORS middleware
	handler := c.Handler(router)

	// Start API server
	port := ":8080"
	fmt.Printf("API server started at %s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
