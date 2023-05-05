package main

import (
	"fmt"
	"github.com/MasahiroYoshiichi/auth/cognito/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
	router.HandleFunc("/confirm-signup", handlers.ConfirmSignUpHandler).Methods("POST")
	router.HandleFunc("/signin", handlers.SignInHandler).Methods("POST")
	router.HandleFunc("/signout", handlers.SignOutHandler).Methods("POST")
	// Start API server
	port := ":8080"
	fmt.Printf("API server started at %s\n", port)
	http.ListenAndServe(port, router)
}
