package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(mode string, trustedProxies []string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.SetTrustedProxies(trustedProxies)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	return r
}
