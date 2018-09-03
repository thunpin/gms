package jwt

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenFromHeader(context *gin.Context) string {
	auth := context.GetHeader("Authorization")
	parts := strings.Split(auth, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}
