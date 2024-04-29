package repositories

import "gorm.io/gorm"

type ServiceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) *ServiceRepository {
	return &ServiceRepository{
		db: db,
	}
}
