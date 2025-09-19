package router

import (
	"backend/internal/errors"
	"backend/internal/models/user"
	"backend/internal/pkg/logger"
	"backend/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, uc user.UserController) *gin.Engine {

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.NewSuccessMessage("Welcome to the Online Judge API"))
	})

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", uc.Register)
		}
	}

	r.NoRoute(func(c *gin.Context) {
		logger.Info().Msg("404 Not Found endpoint hit")
		c.Error(errors.NewNotFoundError("Page Not Found."))
	})

	return r
}
