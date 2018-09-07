package jwt

import (
	"encoding/json"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func FromToken(token string, pointer interface{}, secret string) error {
	jwtToken, err := jwtTokenFromToken(token, secret)
	if err != nil {
		return err
	}

	return fromJwtToken(jwtToken, pointer)
}

func jwtTokenFromToken(token string, secret string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secret), nil
	})
}

func fromJwtToken(jwttoken *jwt.Token, pointer interface{}) error {
	claims, ok := jwttoken.Claims.(jwt.MapClaims)
	if !ok || !jwttoken.Valid {
		return ErrInvalidJwtToken
	}

	currentUnixTime := time.Now().Unix()
	unixTime := int64(claims[timeoutId].(float64))
	if currentUnixTime > unixTime {
		return ErrExpired
	}

	jsonStr := claims[contentId].(string)
	return json.Unmarshal([]byte(jsonStr), pointer)
}
