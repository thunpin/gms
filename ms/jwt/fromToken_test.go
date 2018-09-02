package jwt

import (
	"testing"
)

func TestDenyGetFromInvalidToken(t *testing.T) {
	obj := newObj()
	err := FromToken("", &obj, "")
	if err == nil {
		t.Fail()
	}
}

func TestDenyGetFromExpiredToken(t *testing.T) {
	obj := newObj()
	token, _ := CreateToken(newObj(), "1234567890", -1)
	err := FromToken(token, &obj, "1234567890")
	if err != ErrExpired {
		t.Fail()
	}
}

func TestDenyGetFromInvalidSecret(t *testing.T) {
	obj := newObj()
	token, _ := CreateToken(newObj(), "1234567890", 1)
	err := FromToken(token, &obj, "123456789")
	if err == nil {
		t.Fail()
	}
}

func TestGetFromToken(t *testing.T) {
	obj := newObj()
	token, _ := CreateToken(obj, "1234567890", 1)
	toObj := Obj{}
	err := FromToken(token, &toObj, "1234567890")
	if err != nil {
		t.Fail()
	}

	if obj.Id != toObj.Id && obj.Name != toObj.Name && obj.Date != toObj.Date {
		t.Fail()
	}
}
