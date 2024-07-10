package httpserver

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/outagelab/outagelab/ui"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
)

func (hs *HttpServer) router() *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)

	assetsDir, _ := fs.Sub(ui.Dist, "dist/assets")
	router.StaticFS("/assets", http.FS(assetsDir))

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
			b, _ := ui.Dist.ReadFile("dist/index.html")
			c.Data(http.StatusOK, "text/html; charset=UTF-8", b)
		}
	})

	return router
}
