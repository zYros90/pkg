package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWT struct{}

func New() *JWT { return &JWT{} }

func (j *JWT) GenerateJWTHS256(
	secret string,
	claimsMap map[string]string,
	expire time.Duration,
) (string, error) {
	s := []byte(secret)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	for k, v := range claimsMap {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(expire).Unix()
	return token.SignedString(s)
}

func (j *JWT) ValidateJWTHS256(key []byte, jwtToken string) (map[string]string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claimMap := make(map[string]string)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		for key, val := range claims {
			valStr, ok := val.(string)
			if !ok {
				continue
			}
			claimMap[key] = valStr
		}
	}
	return claimMap, nil
}
