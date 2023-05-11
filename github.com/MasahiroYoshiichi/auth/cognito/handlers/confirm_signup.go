package handlers

import (
	"encoding/json"
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
	"log"
	"net/http"
)

func ConfirmSignUpHandler(w http.ResponseWriter, r *http.Request) {

	// AWS設定ファイル読み込み
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("設定ファイルの読み込みに失敗しました。: %v\n", err)
		http.Error(w, "サーバー内部エラー: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// username Cookie　取得
	var confirmUser string
	usernameCookie, err := r.Cookie("username")
	if err != nil {
		log.Printf("UsernameのCookieが取得できませんでした。: %v\n", err)
		http.Error(w, "サーバー内部エラー: "+err.Error(), http.StatusBadRequest)
		return
	}

	// username Cookie　変数へ格納
	confirmUser = usernameCookie.Value
	log.Printf("UsernameのCookieを表示します。: %s\n", usernameCookie)

	// リクエスト処理
	var confirmCode models.ConfirmCode
	err = json.NewDecoder(r.Body).Decode(&confirmCode)
	if err != nil {
		log.Printf("リクエストボディの読み込みに失敗しました: %v\n", err)
		http.Error(w, "不正なリクエスト: "+err.Error(), http.StatusBadRequest)
		return
	}

	// AWS認証情報を元にServiceを作成
	confirmSignUpService := services.NewConfirmSignUpService(cfg)

	// AWSCognitoサインアップ検証
	err = confirmSignUpService.ConfirmSignUp(confirmUser, confirmCode)
	if err != nil {
		log.Printf("検証に失敗しました: %s\n", err)
		http.Error(w, "サーバー内部エラー: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// httpステータス返却
	w.WriteHeader(http.StatusOK)
}
