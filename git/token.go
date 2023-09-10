package git

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	PrivateKeyPath = "./private-key.pem"
)

func GetToken() (string, error) {

	expTime := time.Now().Add(5 * time.Minute)
	jwtClaims := jwt.RegisteredClaims{
		Issuer:    os.Getenv("APP_ID"),
		ExpiresAt: jwt.NewNumericDate(expTime),
		IssuedAt:  jwt.NewNumericDate(time.Now().Add(-time.Minute)),
	}

	pemKey, _ := os.ReadFile(PrivateKeyPath)
	privateKey, _ := jwt.ParseRSAPrivateKeyFromPEM(pemKey)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, jwtClaims)
	token, err := jwtToken.SignedString(privateKey)

	if err != nil {
		return "", err
	}

	return token, nil
}
