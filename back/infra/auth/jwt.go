package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func NewJWT(subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    os.Getenv("JWT_ISSUER"),
		Subject:   subject,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		NotBefore: jwt.NewNumericDate(time.Now()), // この時間より前のトークンは無効
		IssuedAt:  jwt.NewNumericDate(time.Now()), // トークン生成時間
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.RegisteredClaims)

	if !ok {
		return "", err
	}

	return claims.Subject, nil
}
