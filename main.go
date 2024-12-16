package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/discord"
	"github.com/markbates/goth/providers/github"
	"github.com/xenos112/vertex_backend/db"
	"github.com/xenos112/vertex_backend/middleware"
	"github.com/xenos112/vertex_backend/routes"
	"github.com/xenos112/vertex_backend/routes/oauth"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	var (
		OAUTH_SECRET         = os.Getenv("OAUTH_SECRET")
		DISCORD_CLIENT_ID    = os.Getenv("DISCORD_CLIENT_ID")
		DISCORD_SECRET       = os.Getenv("DISCORD_SECRET")
		DISCORD_REDIRECT_URI = os.Getenv("DISCORD_REDIRECT_URI")
		GITHUB_CLIENT_ID     = os.Getenv("GITHUB_CLIENT_ID")
		GITHUB_SECRET        = os.Getenv("GITHUB_SECRET")
		GITHUB_REDIRECT_URI  = os.Getenv("GITHUB_REDIRECT_URI")
	)

	db.ConnectDB()

	goth.UseProviders(
		discord.New(DISCORD_CLIENT_ID, DISCORD_SECRET, DISCORD_REDIRECT_URI, "identify"),
		github.New(GITHUB_CLIENT_ID, GITHUB_SECRET, GITHUB_REDIRECT_URI),
	)

	store := sessions.NewCookieStore([]byte(OAUTH_SECRET))
	gothic.Store = store
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(middleware.CorsMiddleware())
	router.Use(middleware.Oauth)

	// Routes
	router.GET("/health-check", routes.HealthCheck)
	router.GET("/user/:tag", routes.GetUserByTag)
	auth := router.Group("/auth")
	authenticated := router.Group("/authenticated")
	authenticated.Use(middleware.Auth())
	authenticated.GET("/me", routes.Me)
	authenticated.GET("/who-to-follow", routes.WhoToFollow)
	auth.POST("/login", routes.Login)
	auth.POST("/register", routes.Register)
	auth.GET("/:provider", oauth.Provider)
	auth.GET("/:provider/callback", oauth.CallBack)

	router.Run(":8080")
}
