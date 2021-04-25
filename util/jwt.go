package util

import (
	"github.com/dgrijalva/jwt-go"
)

// ParseToken 解析 token
func ParseToken(token string, secret string) (*jwt.StandardClaims, error) {

	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*jwt.StandardClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}
