package handlers

import (
	"github.com/MasahiroYoshiichi/auth/cognito/cognitotoken"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
	"net/http"
)

func SignOutHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	accessToken, err := cognitotoken.AccessTokenFromContext(r.Context())
	if err != nil {
		http.Error(w, "ヘッダーに認証情報がありません。", http.StatusUnauthorized)
		return
	}

	signOutService := services.NewSignOutService(cfg)
	err = signOutService.SignOut(accessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
