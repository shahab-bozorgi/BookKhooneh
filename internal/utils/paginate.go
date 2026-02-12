package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type PaginationResult[T any] struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
	Data  []T `json:"data"`
}

func Paginate[T any](c *gin.Context, db *gorm.DB) (PaginationResult[T], error) {
	var result PaginationResult[T]

	page := c.DefaultQuery("page", "1")
	size := c.DefaultQuery("size", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}
	sizeInt, err := strconv.Atoi(size)
	if err != nil || sizeInt <= 0 {
		sizeInt = 10
	}

	offset := (pageInt - 1) * sizeInt

	var data []T
	tx := db.Limit(sizeInt).Offset(offset)
	if err := tx.Find(&data).Error; err != nil {
		return result, err
	}

	var total int64
	db.Model(new(T)).Count(&total)

	result = PaginationResult[T]{
		Page:  pageInt,
		Size:  sizeInt,
		Total: int(total),
		Data:  data,
	}

	return result, nil
}
