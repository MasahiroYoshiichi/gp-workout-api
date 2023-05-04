package handlers

import (
	"fmt"
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
)

func SignInHandler(cfg *config.Config, signinInfo models.AuthInfo) {
	signInService := services.NewSignInService(cfg)
	initiateAuthOutput, err := signInService.SignIn(signinInfo)
	if err != nil {
		fmt.Println("Error signing in:", err)
		return
	}
	fmt.Println("SignIn result:", initiateAuthOutput)
}
