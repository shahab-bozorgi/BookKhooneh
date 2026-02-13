package application

import (
	"BookKhoone/internal/domain"
	"gorm.io/gorm"
)

func GetUserService(db *gorm.DB, username string) (domain.User, error) {
	var user domain.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func GetAllUsersService(db *gorm.DB) ([]domain.User, error) {
	var users []domain.User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
