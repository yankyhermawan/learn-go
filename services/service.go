package services

import (
	"github.com/yankyhermawan/learn-go/models"
	"gorm.io/gorm"
)

func FindByID(db *gorm.DB, productID uint64) (*models.Product, error) {
	var product models.Product
	res := db.First(&product, productID)
	if res.Error != nil {
		return nil, res.Error
	}
	return &product, nil
}

func getAll(db *gorm.DB) (*models.Product, error) {
	var product models.Product
	res := db.Find(&product)
	if res.Error != nil {
		return nil, res.Error
	}
	return &product, nil
}

func CreateUser(db *gorm.DB, product *models.Product) error {
	result := db.Create(product)
	return result.Error
}
