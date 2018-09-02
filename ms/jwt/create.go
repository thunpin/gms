package jwt

import (
	"encoding/json"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(
	content interface{},
	secret string,
	signedMethod jwt.SigningMethod,
	days int) (string, error) {

	json := toJSON(content)
	token := jwt.NewWithClaims(signedMethod, jwt.MapClaims{
		contentId: json,
		timeoutId: time.Now().AddDate(0, 0, days).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	return token.SignedString([]byte(secret))
}

func toJSON(content interface{}) string {
	bytes, _ := json.Marshal(content)
	return string(bytes)
}
