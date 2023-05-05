package models

type AuthInfo struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phoneNumber"`
	Password         string `json:"password"`
	ConfirmationCode string `json:"confirmationCode"`
	AccessToken      string `json:"accessToken"`
}
