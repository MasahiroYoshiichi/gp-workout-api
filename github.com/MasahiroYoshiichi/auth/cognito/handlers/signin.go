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

	var signinInfo models.SignInInfo
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

	idToken := initiateAuthOutput.AuthenticationResult.IdToken
	if idToken == nil {
		http.Error(w, "ID token is missing", http.StatusInternalServerError)
		return
	}
	//tokenManager := customtoken.NewTokenManager(cfg.JWtSecret)
	//token, err := tokenManager.GenerateToken(signinInfo.Username)
	//if err != nil {
	//	http.Error(w, "トークンが作成できませんでした。", http.StatusInternalServerError)
	//	return
	//}

	response := struct {
		Token string `json:"token"`
	}{
		Token: *idToken,
	}

	json.NewEncoder(w).Encode(response)
}
