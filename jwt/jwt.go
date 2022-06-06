package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWTHS256(secret string, claimsMap map[string]string, expire time.Duration) (string, error) {
	s := []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range claimsMap {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(expire).Unix()
	return token.SignedString(s)
}
