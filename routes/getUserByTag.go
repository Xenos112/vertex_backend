package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xenos112/vertex_backend/db"
	"gorm.io/gorm"
)

func GetUserByTag(c *gin.Context) {
	var user db.User

	tag := c.Param("tag")

	if tag == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "tag is required",
		})
		return
	}

	err := db.DB.Raw("SELECT * FROM users WHERE tag = ?", tag).First(&user).Error
	if err == gorm.ErrRecordNotFound {
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
