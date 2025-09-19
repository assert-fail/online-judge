package middleware

import (
	"backend/internal/models/errors"
	"backend/internal/models/response"
	"backend/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errorNum := len(c.Errors)
		if errorNum > 0 {
			logger.Error().
				Err(c.Errors[0]).
				Int("error_count", errorNum).
				Msg("Handling error in middleware")

			switch err := c.Errors[0].Err.(type) {
			case *errors.AppError:
				c.JSON(
					err.Code,
					response.NewErrorData(
						err.Message,
						err.Details,
					),
				)
			case validator.ValidationErrors:
				c.JSON(
					http.StatusBadRequest,
					response.NewErrorData(
						"Validation failed",
						formatValidationErrors(err),
					),
				)
			default:
				// 生产环境中隐藏具体错误细节
				if gin.Mode() == gin.ReleaseMode {
					c.JSON(
						http.StatusInternalServerError,
						response.NewErrorMessage("Internal server error"),
					)
				} else {
					c.JSON(
						http.StatusInternalServerError,
						response.NewErrorData(
							"Internal server error",
							err.Error(),
						),
					)
				}
			}
		}
	}
}

func formatValidationErrors(errs validator.ValidationErrors) []string {
	var errors []string
	for _, err := range errs {
		errors = append(errors, err.Field()+": "+err.Tag())
	}
	return errors
}
