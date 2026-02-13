package application

import (
	utils2 "BookKhoone/infrastructure/utils"
	"BookKhoone/internal/domain"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username, email, password string) (*domain.User, error) {
	hashedPassword, err := utils2.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
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
	return utils2.GenerateJWT(userID, secret)
}

func LoginUser(db *gorm.DB, username, password string) (*domain.User, error) {
	var user domain.User

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("wrong password")
	}
	return &user, nil
}
