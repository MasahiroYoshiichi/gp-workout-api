package services

import (
	"github.com/MasahiroYoshiichi/auth/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type SignOutService struct {
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
}

func NewSignOutService(cfg *config.Config) *SignOutService {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.AwsRegion)}))
	cognitoClient := cognitoidentityprovider.New(sess)
	return &SignOutService{
		cognitoClient: cognitoClient,
	}
}

func (s *SignOutService) SignOut(accessToken string) error {
	input := &cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: aws.String(accessToken),
	}
	_, err := s.cognitoClient.GlobalSignOut(input)
	return err
}
