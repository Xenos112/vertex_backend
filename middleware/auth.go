package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xenos112/vertex_backend/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("auth_token")
		if err != nil || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		id, err := utils.ParseJWT(token)
		if err != nil || id == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}

		c.Set("id", id)
		c.Next()
	}
}
