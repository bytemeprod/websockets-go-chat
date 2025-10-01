package tokens

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(key []byte, sub string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
	})
	signedToken, err := token.SignedString(key)
	if err != nil {
		fmt.Printf("Failed to sign token: %s", err.Error())
		return "", err
	}
	return signedToken, nil
}

func ValidateJWT(key []byte, tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(tkn *jwt.Token) (any, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", tkn.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
