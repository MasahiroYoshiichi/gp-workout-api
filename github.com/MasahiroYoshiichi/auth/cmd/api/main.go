package main

import (
	"fmt"
	"github.com/MasahiroYoshiichi/auth/cognito/cognitotoken"
	"github.com/MasahiroYoshiichi/auth/cognito/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//　ルーター設定（gorilla使用）
	router := mux.NewRouter()

	// APIエンドポイント
	router.HandleFunc("/signup", handlers.SignUpHandler).Methods("POST")
	router.HandleFunc("/confirm-signup", handlers.ConfirmSignUpHandler).Methods("POST")
	router.HandleFunc("/signin", handlers.SignInHandler).Methods("POST")
	router.HandleFunc("/mfa", handlers.MFAHandler).Methods("POST")
	router.Handle("/signout", cognitotoken.Middleware(http.HandlerFunc(handlers.SignOutHandler))).Methods("POST")

	//　APIサーバー設定
	port := ":8080"
	fmt.Printf("API server started at %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// リバースプロキシを設定認め中断

// Configure CORS 設定する場合
//c := cors.New(cors.Options{
//	AllowedOrigins:   []string{"http://127.0.0.1:5173", "http://localhost:5173"},
//	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
//	AllowedHeaders:   []string{"Authorization", "Content-Type"},
//	AllowCredentials: true,
//})

//ルーターに対してミドルウェアを設定
//handler := c.Handler(router)

//　APIサーバー設定
//	port := ":8080"
//	fmt.Printf("API server started at %s\n", port)
//	log.Fatal(http.ListenAndServe(port, handler))
//}
