package routes

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xenos112/vertex_backend/db"
	"github.com/xenos112/vertex_backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type RegisterJsonData struct {
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context) {
	var data RegisterJsonData
	var oldUser db.User

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if data.Email == "" || data.Password == "" || data.UserName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email and password and username are required"})
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

	err = db.DB.Raw("SELECT * FROM users WHERE email = ?", data.Email).First(&oldUser).Error
	if err == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user with the same email has been found"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	user := db.User{
		ID:         uuid.Must(uuid.NewRandom()),
		UserName:   data.UserName,
		Email:      data.Email,
		Password:   string(hashedPassword),
		Tag:        utils.GenerateRandomTag(),
		Bio:        "",
		Image_ID:   nil,
		Banner_ID:  nil,
		Github_ID:  "",
		Google_ID:  "",
		Discord_ID: "",
	}

	err = db.DB.Exec(`INSERT
			INTO users
			(id, user_name,password,email,tag,created_at,updated_at)
			VALUES (?,?,?,?,?,?,?)`,
		user.ID, user.UserName, user.Password, user.Email, user.Tag, user.CreatedAt, user.UpdatedAt).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	token, err := utils.GenerateJWT(user.ID.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	ctx.SetCookie("auth_token", token, 2592000, "/", "", false, true)
	user.Password = ""

	ctx.JSON(http.StatusOK, gin.H{
		"data":       user,
		"auth_token": token,
	})
}
