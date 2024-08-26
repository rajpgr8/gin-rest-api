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
/*
func prepareJWTToken(config *Config) (string, error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(config.PrivateKey.Public())
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(pubBytes)

	accountName := extractAccountName(config.Account)
	userName := strings.ToUpper(config.User)

	issueAtTime := time.Now().UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"iss": fmt.Sprintf("%s.%s.%s", accountName, userName, "SHA256:"+base64.StdEncoding.EncodeToString(hash[:])),
		"sub": fmt.Sprintf("%s.%s", accountName, userName),
		"iat": issueAtTime.Unix(),
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp": issueAtTime.Add(config.JWTExpireTimeout).Unix(),
	})

	tokenString, err := token.SignedString(config.PrivateKey)

	if err != nil {
		return "", err
	}

	return tokenString, err
}
*/
