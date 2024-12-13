package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Allow all origins or specify your frontend domain
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:3000") // Allow all origins, or specify frontend domain e.g., "http://localhost:3000"
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, Accept, X-Requested-With")
		ctx.Header("Access-Control-Allow-Credentials", "true") // Allow cookies to be sent
		ctx.Header("Access-Control-Max-Age", "86400")          // Cache preflight requests for 24 hours

		// If it's a preflight OPTIONS request, return immediately
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
			return
		}

		// Continue with the next handler
		ctx.Next()
	}
}
