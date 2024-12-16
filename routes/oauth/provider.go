package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
)

func Provider(c *gin.Context) {
	provider := c.Param("provider")
	if _, err := goth.GetProvider(provider); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid provider"})
		return
	}

	gothic.GetProviderName = func(*http.Request) (string, error) {
		return provider, nil
	}
	gothic.BeginAuthHandler(c.Writer, c.Request)
}
