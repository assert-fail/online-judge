package user

import (
	"backend/internal/errors"
	"backend/internal/models/user/request"
	"backend/internal/pkg/logger"
	"backend/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
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
		log.Warn().Err(err).Msg("Failed to bind request body")
		c.Error(err)
		return
	}

	if err := uc.service.Register(&req); err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			log.Warn().Err(appErr).Msg("Failed to register user")
		} else {
			log.Error().Err(err).Msg("Database error")
		}
		c.Error(err)
		return
	}

	c.JSON(
		http.StatusOK,
		response.NewSuccessMessage("User registered successfully"),
	)
}

func (uc *userController) Login(c *gin.Context) {
	log := logger.WithRequest(c)
	log.Info().Msg("User login endpoint hit")

	var req request.LoginBody
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		log.Warn().Err(err).Msg("Failed to bind request body")
		c.Error(err)
		return
	}

	token, err := uc.service.Login(&req)
	if err != nil {
		if appErr, ok := err.(*errors.AppError); ok {
			log.Warn().Err(appErr).Msg("Failed to login user")
		} else {
			log.Error().Err(err).Msg("Database error")
		}
		c.Error(err)
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusOK,
		response.NewSuccessMessage(
			"User login successfully",
		),
	)
}
