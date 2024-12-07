package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xenos112/vertex_backend/db"
	"github.com/xenos112/vertex_backend/routes"
)

func init() {
	db.ConnectDB()
}

func main() {
	router := gin.Default()

	// Routes
	router.GET("/health-check", routes.HealthCheck)

	router.Run(":3000")
}
