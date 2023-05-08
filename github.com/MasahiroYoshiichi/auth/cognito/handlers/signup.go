package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	cfg, _ := config.LoadConfig()

	var signupInfo models.SignUpInfo
	err := json.NewDecoder(r.Body).Decode(&signupInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signUpService := services.NewSignUpService(cfg)
	_, err = signUpService.SignUp(signupInfo)
	if err != nil {
		http.Error(w, "サインアップできませんでした。"+err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Username string `json:"username"`
	}{
		Username: signupInfo.Username,
	}

	json.NewEncoder(w).Encode(response)
}
