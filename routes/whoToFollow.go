package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xenos112/vertex_backend/db"
)

func WhoToFollow(c *gin.Context) {
	userID := c.MustGet("id").(string)
	var userTags []db.UserTag
	err := db.DB.Raw("SELECT * FROM user_tags WHERE user_id = ?", userID).Find(&userTags).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}

	type UserSuggestion struct {
		ID           uuid.UUID `json:"id"`
		UserName     string    `json:"user_name"`
		ProfileImage string    `json:"profile_image"`
	}

	var sameTagsUsers []UserSuggestion
	err = db.DB.Raw(`SELECT DISTINCT u.id, u.user_name,
							COALESCE(m.image, '') AS profile_image
					FROM users u
					LEFT JOIN media m ON u.id = m.target_id AND m.type = 'profile_image'
					JOIN user_tags ut ON u.id = ut.user_id
					WHERE ut.tag_id IN (SELECT tag_id FROM user_tags WHERE user_id = ?)
					AND u.id != ?
					GROUP BY u.id, u.user_name, m.image
					HAVING COUNT(DISTINCT ut.tag_id) >= 3
					LIMIT 3`, userID, userID).Scan(&sameTagsUsers).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": sameTagsUsers,
	})
}
