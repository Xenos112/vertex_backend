package routes

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/xenos112/vertex_backend/db"
	"github.com/xenos112/vertex_backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginJsonData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var data LoginJsonData
	var user db.User

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.Email == "" || data.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	_, err := mail.ParseAddress(data.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}

	if len(data.Password) < 8 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password must be at least 8 characters"})
		return
	}

	err = db.DB.Raw("SELECT * FROM users WHERE email = ?", data.Email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not User Found By This Email"})
		return
	}

	if user.Google_ID != "" || user.Discord_ID != "" || user.Github_ID != "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "This Account is Connected with Social Media"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.SetCookie("auth_token", token, 3600, "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"data": "Login Successfully", "auth_token": token})
}
