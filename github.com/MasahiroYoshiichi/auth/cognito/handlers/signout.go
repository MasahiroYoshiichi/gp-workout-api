package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
)

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var signoutInfo models.AuthInfo
	err = json.NewDecoder(r.Body).Decode(&signoutInfo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signOutService := services.NewSignOutService(cfg)
	err = signOutService.SignOut(signoutInfo.AccessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
