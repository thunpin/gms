package jwt

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thunpin/gerrors"
)

func TokenFromHeader(context *gin.Context) (string, error) {
	auth := context.GetHeader("Authorization")
	parts := strings.Split(auth, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", gerrors.Forbidden()
	}

	return parts[1], nil
}
