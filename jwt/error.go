package jwt

import "errors"

var (
	ErrInvalidToken    = errors.New("invalid token")
	ErrInvalidJwtToken = errors.New("invalid JWT token")
	ErrExpired         = errors.New("token expired")
)
