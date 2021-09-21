package controllers

import (
	"gorm.io/gorm"
	"jesus.tn79/aveonline/repository"
)

//
// A Controller response
type Controller struct {
	repo *repository.Repository
}

// NewController function
func NewController(db *gorm.DB) *Controller {
	return &Controller{
		repo: repository.NewRepository(db),
	}
}
