package application

import (
	"BookKhoone/internal/domain"
	"gorm.io/gorm"
)

func CreateBookReviewsService(db *gorm.DB, review domain.Review) (*domain.Review, error) {

	if err := db.Create(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil

}
