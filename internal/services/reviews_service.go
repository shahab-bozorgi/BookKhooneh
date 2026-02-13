package services

import (
	"BookKhoone/internal/models"
	"gorm.io/gorm"
)

func CreateBookReviewsService(db *gorm.DB, review models.Review) (*models.Review, error) {

	if err := db.Create(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil

}
