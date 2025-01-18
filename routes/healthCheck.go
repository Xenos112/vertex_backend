package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Health Check
// @Description Check if the service is up and running.
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{}
// @Router /health-check [get]

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
