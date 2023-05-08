package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
)

func ConfirmSignUpHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var confirmSignupInfo models.ConfirmSignUpInfo
	err = json.NewDecoder(r.Body).Decode(&confirmSignupInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	confirmSignUpService := services.NewConfirmSignUpService(cfg)
	err = confirmSignUpService.ConfirmSignUp(confirmSignupInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
