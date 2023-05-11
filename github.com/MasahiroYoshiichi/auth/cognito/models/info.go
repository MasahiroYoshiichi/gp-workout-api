package models

type SignInInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type MFACode struct {
	MFACode string `json:"MFACode"`
}

type SignUpInfo struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type ConfirmCode struct {
	ConfirmationCode string `json:"confirmationCode"`
}

//type ConfirmUser struct {
//	Username string
//}

type AuthInfo struct {
	AccessToken string `json:"accessToken"`
}
