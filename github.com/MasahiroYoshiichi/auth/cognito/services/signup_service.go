package services

import (
	"github.com/MasahiroYoshiichi/auth/cognito/models"
	"github.com/MasahiroYoshiichi/auth/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// SignUpService サービスの構造体を作成
type SignUpService struct {
	cognitoClient *cognitoidentityprovider.CognitoIdentityProvider
	clientId      string
}

// NewSignUpService サインアップのサービスを作成(引数：AWS設定情報、戻り値：サービスの構造体)
func NewSignUpService(cfg *config.Config) *SignUpService {

	// AWSリージョンとのセッションを確保
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(cfg.AwsRegion)}))

	// Cognitoのクライアントを作成
	cognitoClient := cognitoidentityprovider.New(sess)

	// サービス構造体を作成
	return &SignUpService{
		cognitoClient: cognitoClient,
		clientId:      cfg.ClientId,
	}
}

// SignUp Cognitoへのサインアップを実行(引数；認証情報、戻り値：Cognitoクライアントへの登録情報)
func (s *SignUpService) SignUp(signupInfo models.SignUpInfo) (*cognitoidentityprovider.SignUpOutput, error) {

	// サービスで作成したCognitoクライアントへ登録情報を格納
	signUpInput := &cognitoidentityprovider.SignUpInput{

		// AWS設定情報のクライアントID
		ClientId: aws.String(s.clientId),

		//　登録情報（必須）
		Username: aws.String(signupInfo.Username),
		Password: aws.String(signupInfo.Password),

		//登録情報(追加属性)
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
	//aws.Stringはawsのポインタ型に対して文字列をマッピングするためのもの

	// Cognitoクライアントへの登録を実行
	return s.cognitoClient.SignUp(signUpInput)
}
