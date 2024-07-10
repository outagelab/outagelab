package httpserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func newCorsHandler() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowCredentials = true
	return cors.New(config)
}
