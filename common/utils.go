package common

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateToken() string {
	code := uuid.New().String()
	code = strings.Replace(code, "-", "", -1)
	return code
}
