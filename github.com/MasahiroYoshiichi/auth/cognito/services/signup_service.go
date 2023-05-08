package services

import (
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type SignUpService struct {
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	clientId      string
}

func NewSignUpService(cfg *config.Config) *SignUpService {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.AwsRegion)}))
	cognitoClient := cognitoidentityprovider.New(sess)
	return &SignUpService{
		cognitoClient: cognitoClient,
		clientId:      cfg.ClientId,
	}
}

func (s *SignUpService) SignUp(signupInfo models.SignUpInfo) (*cognitoidentityprovider.SignUpOutput, error) {
	signUpInput := &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(s.clientId),
		Username: aws.String(signupInfo.Username),
		Password: aws.String(signupInfo.Password),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String(signupInfo.Email),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String(signupInfo.PhoneNumber),
			},
		},
	}

	return s.cognitoClient.SignUp(signUpInput)
}
