package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xenos112/vertex_backend/db"
	"github.com/xenos112/vertex_backend/middleware"
	"github.com/xenos112/vertex_backend/routes"
)

func init() {
	db.ConnectDB()
}

func main() {
	router := gin.Default()

	// Routes
	router.GET("/health-check", routes.HealthCheck)
	router.GET("/user/:tag", routes.GetUserByTag)
	auth := router.Group("/auth")
	authenticated := router.Group("/authenticated")
	authenticated.Use(middleware.Auth())
	authenticated.GET("/me", routes.Me)
	auth.POST("/login", routes.Login)
	auth.POST("/register", routes.Register)

	router.Run(":3000")
}
