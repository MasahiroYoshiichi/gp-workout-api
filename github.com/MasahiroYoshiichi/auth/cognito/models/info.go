package models

type SignInInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpInfo struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type ConfirmSignUpInfo struct {
	Username         string `json:"username"`
	ConfirmationCode string `json:"confirmationCode"`
}

type AuthInfo struct {
	AccessToken string `json:"accessToken"`
}
