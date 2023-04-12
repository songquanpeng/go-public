package common

import (
	"github.com/google/uuid"
	"os/user"
)

func GenerateToken() string {
	code := uuid.New().String()
	return code
}

func Token2Bytes(token string) []byte {
	id, err := uuid.Parse(token)
	if err != nil {
		return []byte(token)
	}
	return id[:]
}

func Bytes2Token(bytes []byte) string {
	id, err := uuid.FromBytes(bytes)
	if err != nil {
		return string(bytes)
	}
	return id.String()
}

func GetHomeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
}
