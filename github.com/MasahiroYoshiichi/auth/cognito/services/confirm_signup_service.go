package services

import (
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// ConfirmSignUpService サービスの構造体を作成
type ConfirmSignUpService struct {
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	clientId      string
}

// NewConfirmSignUpService  サインアップ検証のサービスを作成(引数：AWS設定情報、戻り値：サービスの構造体)
func NewConfirmSignUpService(cfg *config.Config) *ConfirmSignUpService {

	// AWSリージョンとのセッションを確保
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.AwsRegion)}))

	// Cognitoのクライアントを作成
	cognitoClient := cognitoidentityprovider.New(sess)

	// サービス構造体を作成
	return &ConfirmSignUpService{
		cognitoClient: cognitoClient,
		clientId:      cfg.ClientId,
	}
}

// ConfirmSignUp Cognitoへのサインアップ検証を実行(引数；認証情報)
func (s *ConfirmSignUpService) ConfirmSignUp(confirmUser string, confirmCode models.ConfirmCode) error {

	// サービスで作成したCognitoクライアントへ検証情報を格納
	confirmSignInInput := &cognitoidentityprovider.ConfirmSignUpInput{

		// AWS設定情報のクライアントID
		ClientId: aws.String(s.clientId),

		// 検証情報
		Username:         aws.String(confirmUser),
		ConfirmationCode: aws.String(confirmCode.ConfirmationCode),
	}

	// Cognitoクライアントへの検証を実行
	_, err := s.cognitoClient.ConfirmSignUp(confirmSignInInput)
	return err
}
