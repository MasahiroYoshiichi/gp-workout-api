package handlers

import (
	"encoding/json"
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
	"log"
	"net/http"
	"time"
)

func SignInHandler(w http.ResponseWriter, r *http.Request) {

	// AWS設定ファイル読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("設定ファイルの読み込みに失敗しました。:", err)
		http.Error(w, "サーバー内部エラー", http.StatusInternalServerError)
		return
	}

	// リクエスト処理
	var signinInfo models.SignInInfo
	err = json.NewDecoder(r.Body).Decode(&signinInfo)
	if err != nil {
		log.Printf("リクエストボディの読み込みに失敗しました: %v\n", err)
		http.Error(w, "不正なリクエスト", http.StatusBadRequest)
		return
	}

	// AWS認証情報を元にServiceを作成
	signInService := services.NewSignInService(cfg)

	// AWSCognitoサインイン
	initiateAuthOutput, err := signInService.SignIn(signinInfo)
	if err != nil {
		log.Printf("認証に失敗しました。: %v\n", err)
		http.Error(w, "サーバー内部エラー："+err.Error(), http.StatusInternalServerError)
		return
	}

	// email Cookie　格納
	http.SetCookie(w, &http.Cookie{
		Name:     "email",
		Value:    signinInfo.Email,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(30 * time.Minute),
	})
	log.Printf("Cookieを設定(email): %s\n", signinInfo.Email)

	// Session Cookie　格納
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    *initiateAuthOutput.Session,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(30 * time.Minute),
	})
	log.Printf("Cookieを設定(session): %s\n", *initiateAuthOutput.Session)

	// httpステータス返却
	w.WriteHeader(http.StatusOK)
}
