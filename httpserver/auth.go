package httpserver

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/outagelab/outagelab/models"
)

func (hs *HttpServer) authMiddleware(ctx *gin.Context) {
	defer ctx.Next()

	auth := ctx.GetHeader("Authorization")
	authSegments := strings.Fields(auth)
	if len(authSegments) != 2 || strings.ToLower(authSegments[0]) != "bearer" {
		return
	}

	authToken := authSegments[1]
	claims, err := hs.authService.ValidateToken(authToken)
	if err != nil {
		// TODO: log
	}
	if claims == nil {
		return
	}

	setAuthClaims(ctx, claims)
}

func setAuthClaims(ctx *gin.Context, claims *models.AuthClaims) {
	ctx.Set("custom_auth_claims", claims)
}

func getAuthClaims(ctx *gin.Context) *models.AuthClaims {
	claims, exist := ctx.Get("custom_auth_claims")
	if !exist {
		return nil
	}

	return claims.(*models.AuthClaims)
}
