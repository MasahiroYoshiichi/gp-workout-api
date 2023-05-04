package handlers

import (
	"fmt"
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
)

func SignUpHandler(cfg *config.Config, signupInfo models.AuthInfo) {
	signUpService := services.NewSignUpService(cfg)
	signUpOutput, err := signUpService.SignUp(signupInfo)
	if err != nil {
		fmt.Println("サインアップできませんでした。", err)
		return
	}
	fmt.Println("SignUp result:", signUpOutput)
}
