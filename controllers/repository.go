package controllers

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func Controllers(db *gorm.DB) Repository {
	return Repository {
		db: db,
	}
}
