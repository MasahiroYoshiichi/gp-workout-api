package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var signinInfo models.AuthInfo
	err = json.NewDecoder(r.Body).Decode(&signinInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signInService := services.NewSignInService(cfg)
	initiateAuthOutput, err := signInService.SignIn(signinInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(initiateAuthOutput)
}
