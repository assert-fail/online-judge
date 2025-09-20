package middleware

import (
	"backend/internal/pkg/database"
	"backend/internal/pkg/utils"
	"backend/internal/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || len(token) < 7 || token[:7] != "Bearer " {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.NewErrorMessage("Missing or invalid token"),
			)
			return
		}

		token = token[7:]

		claim, err := utils.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.NewErrorMessage("Missing or invalid token"),
			)
			return
		}

		foundToken, err := database.RDBInstance.Get(c, "token"+claim.Username).Result()
		if err != nil || foundToken != token {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				response.NewErrorMessage("Missing or invalid token"),
			)
			return
		}

		c.Set("userID", claim.UserID)

		c.Next()
	}
}
