package user

import (
	"backend/internal/models/user/request"
	"backend/internal/pkg/logger"
	"backend/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(c *gin.Context)
}

type userController struct {
	service UserService
}

func New(service UserService) UserController {
	return &userController{service: service}
}

func (uc *userController) Register(c *gin.Context) {
	log := logger.WithRequest(c)
	log.Info().Msg("User registration endpoint hit")

	var req request.RegisterBody
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Warn().Msg("Failed to bind request body")
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response.NewSuccessMessage("ok"))
}
