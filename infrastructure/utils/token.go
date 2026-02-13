package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtSecret []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("❌ JWT_SECRET is not set")
	}

	fmt.Println("🔐 JWT_SECRET loaded:", secret)
	jwtSecret = []byte(secret)
}

func main() {
	token, err := GenerateToken(123)
	if err != nil {
		log.Fatal("Error generating token:", err)
	}
	fmt.Println("🔑 Generated JWT Token:", token)

	userID, err := ValidateToken(token)
	if err != nil {
		log.Fatal("Error validating token:", err)
	}
	fmt.Println("✅ Validated UserID:", userID)
}

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["user_id"].(float64)
		if !ok {
			return "", errors.New("user_id not found in token")
		}
		return fmt.Sprintf("%d", int(userID)), nil
	}

	return "", errors.New("invalid token")
}
