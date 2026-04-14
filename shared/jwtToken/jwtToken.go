package jwtToken

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GetSecretKey returns the JWT signing key, reading from env at call time
// (not at package init, which would be before .env is loaded).
func GetSecretKey() []byte {
	return []byte(os.Getenv("ACCESS_TOKEN_SECRET"))
}

func GenerateToken(userID, email, role string) (string, error) {
	// Create claims
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "eduplat",
		},
	}

	// Sign token
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(GetSecretKey())
}
