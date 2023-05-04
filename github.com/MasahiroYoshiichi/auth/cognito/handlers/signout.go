package handlers

import (
	"fmt"
	"github.com/MasahiroYoshiichi/auth/cognito/services"
	"github.com/MasahiroYoshiichi/auth/config"
)

func SignOutHandler(cfg *config.Config, accessToken string) {
	signOutService := services.NewSignOutService(cfg)
	err := signOutService.SignOut(accessToken)
	if err != nil {
		fmt.Println("Error signing out:", err)
		return
	}
	fmt.Println("SignOut succeeded")
}
