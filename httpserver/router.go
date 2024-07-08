package httpserver

import (
	"strings"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

func (hs *HttpServer) router() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	router.Static("/assets", "./dist/assets")

	router.Use(newCorsHandler())
	router.Use(gintrace.Middleware("outagelab"))
	router.Use(hs.authMiddleware)

	router.POST("/api/login", hs.handleLogin)

	router.GET("/api/account", hs.handleGetAccount)
	router.PUT("/api/account", hs.handleUpdateAccount)

	router.POST("/api/datapage", hs.handleGenerateDatapageDeprecated)
	router.POST("/api/v1/datapage", hs.handleGenerateDatapage)

	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path + "/"
		if c.Request.Method != "GET" || strings.HasPrefix(path, "/api/") || strings.HasPrefix(path, "/assets/") {
			c.AbortWithStatus(404)
		} else {
			c.File("./dist/index.html")
		}
	})

	return router
}
