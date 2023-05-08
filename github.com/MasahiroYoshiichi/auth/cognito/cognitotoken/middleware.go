package cognitotoken

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/MasahiroYoshiichi/auth/config" // Add this import
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
)

type contextKey string

const (
	accessTokenContextKey contextKey = "accessToken"
	usernameContextKey    contextKey = "username"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg, err := config.LoadConfig()
		if err != nil {
			http.Error(w, "設定ファイルがありません。", http.StatusInternalServerError)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "ヘッダーに認証情報がありません。", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		set, err := fetchCognitoJWKSet(cfg.UserPoolId, cfg.AwsRegion)
		if err != nil {
			http.Error(w, "JWKKeyが取得できませんでした。", http.StatusInternalServerError)
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			kid, ok := token.Header["kid"].(string)
			if !ok {
				return nil, errors.New("キーIDが取得できませんでした。")
			}

			keyCount := set.Len()
			for i := 0; i < keyCount; i++ {
				keyIface, _ := set.Get(i)
				if key, ok := keyIface.(jwk.Key); ok {
					if key.KeyID() == kid {
						if rsaKey, ok := key.(jwk.RSAPublicKey); ok {
							pubKey, err := rsaKey.PublicKey()
							if err != nil {
								return nil, err
							}
							return pubKey, nil
						}
					}
				}
			}

			return nil, errors.New("JWKがセットされていません。")
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		username, ok := claims["cognito:username"].(string)
		if !ok {
			http.Error(w, "No username in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), usernameContextKey, username)
		ctx = context.WithValue(ctx, accessTokenContextKey, tokenStr) // Add this line to store the access token in context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AWS cognito　JWLKEYの取得
func fetchCognitoJWKSet(userPoolID, region string) (jwk.Set, error) {
	jwkURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)
	resp, err := http.Get(jwkURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	keySet, err := jwk.ParseReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return keySet, nil
}

func UsernameFromContext(ctx context.Context) (string, error) {
	username, ok := ctx.Value(usernameContextKey).(string)
	if !ok {
		return "", errors.New("no username in context")
	}

	return username, nil
}

func AccessTokenFromContext(ctx context.Context) (string, error) {
	accessToken, ok := ctx.Value(accessTokenContextKey).(string)
	if !ok {
		return "", errors.New("no access token in context")
	}

	return accessToken, nil
}
