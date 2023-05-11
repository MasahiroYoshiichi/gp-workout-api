package services

import (
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func (s *SignInService) CompleteMFA(mfaEmail string, mfaSession string, mfaCode models.MFACode) (*cognitoidentityprovider.AuthenticationResultType, error) {
	input := &cognitoidentityprovider.RespondToAuthChallengeInput{
		ClientId:      aws.String(s.clientId),
		ChallengeName: aws.String(cognitoidentityprovider.ChallengeNameTypeSmsMfa),
		Session:       aws.String(mfaEmail),
		ChallengeResponses: map[string]*string{
			"USERNAME":     aws.String(mfaSession),
			"SMS_MFA_CODE": aws.String(mfaCode.MFACode),
		},
	}

	res, err := s.cognitoClient.RespondToAuthChallenge(input)
	if err != nil {
		return nil, err
	}

	return res.AuthenticationResult, nil
}
