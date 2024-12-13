package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xenos112/vertex_backend/db"
	"gorm.io/gorm"
)

type UserWithProfileImage struct {
	UserID       uuid.UUID `json:"user_id"`
	UserName     string    `json:"user_name"`
	Email        string    `json:"email"`
	Tag          string    `json:"tag"`
	Bio          string    `json:"bio"`
	ProfileImage *string   `json:"profile_image"` // Nullable profile image
}

func Me(c *gin.Context) {
	var user UserWithProfileImage
	userID := c.MustGet("id").(string)

	err := db.DB.Raw(`SELECT 
    u.id AS user_id,
    u.user_name,
    u.email,
    u.tag,
    u.bio,
    m.image AS profile_image
			FROM 
    users u
		LEFT JOIN 
    media m ON m.target_id = u.id
		WHERE 
    u.id = ? `, userID).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		c.SetCookie("auth_token", "", -1, "/", "", false, true)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
