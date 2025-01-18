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

type errorResponse struct {
	Error string `json:"error"`
}

// @Summary Login
// @Description Login to the service
// @Tags login
// @Accept  json
// @Produce  json
// @Param login body LoginJsonData true "Login Data"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Router /login [post]

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

	ctx.SetCookie("auth_token", token, 2592000, "/", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"data": user, "auth_token": token})
}
