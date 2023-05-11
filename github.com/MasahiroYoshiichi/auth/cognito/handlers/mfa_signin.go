package handlers

import (
	"encoding/json"
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
	"log"
	"net/http"
)

func MFAHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//email Cookie　取得
	var mfaEmail string
	emailCookie, err := r.Cookie("email")
	log.Printf("email Cookie: %s\n", emailCookie)
	if err != nil {
		http.Error(w, "EmailのCookieが取得できませんでした。", http.StatusBadRequest)
	}

	//session Cookie 取得
	var mfaSession string
	sessionCookie, err := r.Cookie("session")
	log.Printf("session Cookie: %s\n", sessionCookie)
	if err != nil {
		http.Error(w, "SessionのCookieが取得できませんでした。", http.StatusBadRequest)
	}

	//MFACode 取得
	var mfaCode models.MFACode
	err = json.NewDecoder(r.Body).Decode(&mfaCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signInService := services.NewSignInService(cfg)
	authenticationResult, err := signInService.CompleteMFA(mfaEmail, mfaSession, mfaCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := struct {
		Token string `json:"token"`
	}{
		Token: *authenticationResult.IdToken,
	}

	json.NewEncoder(w).Encode(response)
}
