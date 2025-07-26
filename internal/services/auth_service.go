package services

import (
	"BookKhoone/internal/models"
	"BookKhoone/internal/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username, email, password string) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
	}

	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GenerateUserToken(userID uint, secret string) (string, error) {
	return utils.GenerateJWT(userID, secret)
}

func LoginUser(db *gorm.DB, username, password string) (*models.User, error) {
	var user models.User

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("wrong password")
	}
	return &user, nil
}
