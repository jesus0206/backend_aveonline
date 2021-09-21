package repository

import (
	"gorm.io/gorm"
	"jesus.tn79/aveonline/model"
	// "jesus.tn79/ficha_electronica/models"
)

// A Repository connection db
type Repository struct {
	db *gorm.DB
}

// NewRepository exported
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// IDataBase interface
type IRepository interface {
	GetMedicamentos() ([]*model.Medicamento, error)
	CreatePromociones(model.Promocion) (*string, error)
	GetPromociones() ([]*model.Promocion, error)
	CreateMedicamento(model.Medicamento) (*string, error)
	GetFacturas(model.FacturaDto) ([]*model.Promocion, error)
	CreateFactura(model.FacturaCreateDto) (*string, error)
	SimularFactura(model.FacturaSimularDto) (*float32, error)
}
