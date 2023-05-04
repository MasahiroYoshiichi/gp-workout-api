package main

import (
	"fmt"
	"github.com/MasahiroYoshiichi/auth/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func main() {
	cfg, err := config.LoadCpnfig()
	if err != nil {
		fmt.Println("設定が読み込めませんでした。", err)
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.AwsRegion),
	})
	if err != nil {
		fmt.Println("セッションが作成できませんでした。", err)
		return
	}

	cognitoClient := cognitoidentityprovider.New(sess)

	//userPoolID := cfg.UserPoolId
	clientID := cfg.ClientId

	// SignUp
	signUpOutput, err := cognitoClient.SignUp(&cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(clientID),
		Username: aws.String("username"),
		Password: aws.String("password"),
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: aws.String("email@example.com"),
			},
			{
				Name:  aws.String("phone_number"),
				Value: aws.String("+1234567890"),
			},
		},
	})
	if err != nil {
		fmt.Println("Error signing up:", err)
		return
	}
	fmt.Println("SignUp result:", signUpOutput)

	// ConfirmSignUp
	_, err = cognitoClient.ConfirmSignUp(&cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         aws.String(clientID),
		Username:         aws.String("username"),
		ConfirmationCode: aws.String("confirmation_code"),
	})
	if err != nil {
		fmt.Println("Error confirming sign up:", err)
		return
	}
	fmt.Println("ConfirmSignUp succeeded")

	// SignIn
	initiateAuthOutput, err := cognitoClient.InitiateAuth(&cognitoidentityprovider.InitiateAuthInput{
		ClientId: aws.String(clientID),
		AuthFlow: aws.String(cognitoidentityprovider.AuthFlowTypeUserPasswordAuth),
		AuthParameters: map[string]*string{
			"USERNAME": aws.String("username"),
			"PASSWORD": aws.String("password"),
		},
	})
	if err != nil {
		fmt.Println("Error signing in:", err)
		return
	}
	fmt.Println("SignIn result:", initiateAuthOutput)

	// SignOut
	_, err = cognitoClient.GlobalSignOut(&cognitoidentityprovider.GlobalSignOutInput{
		AccessToken: initiateAuthOutput.AuthenticationResult.AccessToken,
	})
	if err != nil {
		fmt.Println("Error signing out:", err)
		return
	}
	fmt.Println("SignOut succeeded")
}
