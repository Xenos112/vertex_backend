package oauth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/xenos112/vertex_backend/db"
)

func createUser(user goth.User, provider string) (*db.User, error) {
	prevUser := db.User{}
	newUser := db.User{}
	imageProfile := db.Media{}

	switch provider {
	case "discord":
		db.DB.Raw("SELECT * FROM users WHERE discord_id = ?", user.UserID).Scan(&prevUser)
		break
	case "github":
		db.DB.Raw("SELECT * FROM users WHERE github_id = ?", user.UserID).Scan(&prevUser)
		break
	case "google":
		db.DB.Raw("SELECT * FROM users WHERE github_id = ?", user.UserID).Scan(&prevUser)
		break
	}

	if prevUser.ID != uuid.Nil {
		return &prevUser, nil
	}

	switch provider {
	case "discord":
		db.DB.Exec("INSERT INTO users (id,user_name,discord_id) VALUES (?,?,?) RETURNING *", uuid.Must(uuid.NewUUID()), user.NickName, user.UserID).First(&newUser)
		break
	case "github":
		db.DB.Exec("INSERT INTO users (id,user_name,github_id) VALUES (?,?,?) RETURNING *", uuid.Must(uuid.NewUUID()), user.NickName, user.UserID).First(&newUser)
		break
	case "google":
		db.DB.Exec("INSERT INTO users (id,user_name,google_id) VALUES (?,?,?) RETURNING *", uuid.Must(uuid.NewUUID()), user.NickName, user.UserID).First(&newUser)
		break
	}

	err := db.DB.Exec("INSERT INTO media (id, target_id, type, image) VALUES (?,?,?,?) RETURNING *",
		uuid.Must(uuid.NewUUID()), newUser.ID, "image", user.AvatarURL).First(&imageProfile).Error
	if err != nil {
		return nil, err
	}

	newUser.Image_ID = (&imageProfile.ID)
	err = db.DB.Exec("UPDATE users SET image_id = ? WHERE id = ?", imageProfile.ID, newUser.ID).Error
	if err != nil {
		return nil, err
	}

	fmt.Print(newUser)

	return &newUser, nil
}

func CallBack(c *gin.Context) {
	provider := c.Param("provider")
	if _, err := goth.GetProvider(provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provider"})
		return
	}

	gothic.GetProviderName = func(*http.Request) (string, error) {
		return provider, nil
	}

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, err := createUser(user, provider)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed To Create User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": newUser,
	})
}
