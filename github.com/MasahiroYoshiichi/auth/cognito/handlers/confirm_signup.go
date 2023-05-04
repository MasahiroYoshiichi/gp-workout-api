package handlers

import (
	"fmt"
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
)

func ConfirmSignUpHandler(cfg *config.Config, confirmSignupInfo models.AuthInfo) {
	confirmSignUpService := services.NewConfirmSignUpService(cfg)
	err := confirmSignUpService.ConfirmSignUp(confirmSignupInfo)
	if err != nil {
		fmt.Println("Error confirming sign up:", err)
		return
	}
	fmt.Println("ConfirmSignUp succeeded")
}
