package services

import (
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type ConfirmSignUpService struct {
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	clientId      string
}

func NewConfirmSignUpService(cfg *config.Config) *ConfirmSignUpService {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.AwsRegion)}))
	cognitoClient := cognitoidentityprovider.New(sess)
	return &ConfirmSignUpService{
		cognitoClient: cognitoClient,
		clientId:      cfg.ClientId,
	}
}

func (s *ConfirmSignUpService) ConfirmSignUp(confirmSignupInfo models.AuthInfo) error {
	input := &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(s.clientId),
		Username:         aws.String(confirmSignupInfo.Username),
		ConfirmationCode: aws.String(confirmSignupInfo.ConfirmationCode),
	}
	_, err := s.cognitoClient.ConfirmSignUp(input)
	return err
}
