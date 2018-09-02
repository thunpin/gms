package jwt

import (
	"testing"
)

func TestCreate(t *testing.T) {
	token, err := CreateToken(newObj(), "1234567890", 1)
	if err != nil {
		t.Fail()
	}

	if len(token) == 0 {
		t.Fail()
	}
}
