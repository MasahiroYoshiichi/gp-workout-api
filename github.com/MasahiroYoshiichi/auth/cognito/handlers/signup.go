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

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// AWS設定ファイル読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("設定ファイルの読み込みに失敗しました。: %v\n", err)
		http.Error(w, "サーバー内部エラー："+err.Error(), http.StatusInternalServerError)
		return
	}

	// リクエスト処理
	var signupInfo models.SignUpInfo
	err = json.NewDecoder(r.Body).Decode(&signupInfo)
	if err != nil {
		log.Printf("リクエストボディの読み込みに失敗しました。: %v\n", err)
		http.Error(w, "不正なリクエスト："+err.Error(), http.StatusBadRequest)
		return
	}

	// AWS認証情報を元にServiceを作成
	signUpService := services.NewSignUpService(cfg)

	// AWSCognitoサインアップ
	_, err = signUpService.SignUp(signupInfo)
	if err != nil {
		log.Printf("登録に失敗しました。: %v\n", err)
		http.Error(w, "サーバー内部エラー："+err.Error(), http.StatusInternalServerError)
		return
	}

	// username Cookie　格納
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    signupInfo.Username,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode, // Change this to None
		Expires:  time.Now().Add(30 * time.Minute),
	})
	log.Printf("Cookieを設定(username): %s\n", signupInfo.Username)

	// httpステータス返却
	w.WriteHeader(http.StatusOK)
}
