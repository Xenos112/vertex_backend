package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateRandomTag() string {
	return strings.Split(uuid.New().String(), "-")[0]
}
