package services

import (
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type SignInService struct {
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	clientId      string
}

func NewSignInService(cfg *config.Config) *SignInService {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.AwsRegion)}))
	cognitoClient := cognitoidentityprovider.New(sess)
	return &SignInService{
		cognitoClient: cognitoClient,
		clientId:      cfg.ClientId,
	}
}

func (s *SignInService) SignIn(signinInfo models.AuthInfo) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String(s.clientId),
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeUserPasswordAuth),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String(signinInfo.Username),
			"PASSWORD": aws.String(signinInfo.Password),
		},
	}

	result, err := s.cognitoClient.InitiateAuth(input)
	if err != nil {
		return nil, err
	}

	return result, nil
}