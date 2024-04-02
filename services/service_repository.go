package services

import "gorm.io/gorm"

type ServiceRepository struct {
	db *gorm.DB
}

type IServiceRepository interface {
	NewServiceRepository(db *gorm.DB) ServiceRepository
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return ServiceRepository{
		db: db,
	}
}
