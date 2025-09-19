package router

import (
	"backend/internal/models/errors"
	"backend/internal/models/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.NewSuccessMessage("Welcome to the Online Judge API"))
	})

	r.NoRoute(func(c *gin.Context) {
		c.Error(errors.NewNotFoundError("Page Not Found."))
	})

	return r
}
