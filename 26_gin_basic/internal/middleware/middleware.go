package middleware

import (
	"errors"
	"gin-basic/pkg/response"
	jwttoken "gin-basic/pkg/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthToken(ctx *gin.Context) {
	type AuthHeaders struct {
		Authorization string `header:"Authorization" binding:"required"`
		UserAgent     string `header:"User-Agent"`
		ContentType   string `header:"Content-Type" binding:"required,contains=application/json"`
	}
	var authHeaders AuthHeaders
	err := ctx.ShouldBindHeader(&authHeaders)
	if err != nil {
		ctx.AbortWithStatusJSON(response.Error(http.StatusBadRequest, err))
		return
	}

	if !strings.HasPrefix(authHeaders.Authorization, "Bearer ") {
		err := errors.New("bearer is required")
		ctx.AbortWithStatusJSON(response.Error(http.StatusUnauthorized, err))
		return
	}

	isValid := jwttoken.IsTokenValid(authHeaders.Authorization[7:], "")
	if !isValid {
		err := errors.New("invalid token")
		ctx.AbortWithStatusJSON(response.Error(http.StatusUnauthorized, err))
		return
	}

	ctx.Next()
}
