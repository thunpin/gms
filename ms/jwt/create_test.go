package jwt

import (
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
)

func TestCreate(t *testing.T) {
	token, err := CreateToken(newObj(), "1234567890", jwt.SigningMethodHS512, 1)
	if err != nil {
		t.Fail()
	}

	if len(token) == 0 {
		t.Fail()
	}
}
