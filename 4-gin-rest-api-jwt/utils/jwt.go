package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "DoNotHardCodeKeyHere"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

// See example to use jwt  https://github.com/snowflakedb/gosnowflake/blob/aa4b7cd4a56e74ac9871595a5e3625e9359ff3ab/auth.go#L482-L508
