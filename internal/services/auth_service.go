package services

import (
	"BookKhoone/internal/models"
	"BookKhoone/internal/utils"
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
