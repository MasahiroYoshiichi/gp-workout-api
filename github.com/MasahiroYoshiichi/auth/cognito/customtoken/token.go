package customtoken

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var ErrInvalidToken = errors.New("invalid token")

type TokenManager struct {
	secret []byte
}

func NewTokenManager(secret string) *TokenManager {
	return &TokenManager{secret: []byte(secret)}
}

func (t *TokenManager) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(t.secret)
}

func (t *TokenManager) VerifyToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return t.secret, nil
	})

	if err != nil {
		return "トークンを変換できませんでした。", err
	}

	if !token.Valid {
		return "トークンが無効です。", ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "トークンの値を変換できませんでした。", ErrInvalidToken
	}

	username, ok := claims["username"].(string)
	if !ok {
		return "値の取得に失敗しました。", ErrInvalidToken
	}

	return username, nil
}
