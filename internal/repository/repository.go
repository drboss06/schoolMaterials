package repository

import (
	"github.com/jmoiron/sqlx"
	"schoolMaterial/internal/models"
)

type MaterialRepository interface {
	CreateMaterial(material models.Material) error
	GetMaterialByUUID(uuid string) (*models.Material, error)
	UpdateMaterial(material models.Material) error
	GetAllMaterials(active bool, materialType, startDate, endDate string) ([]*models.Material, error)
}

type Repository struct {
	MaterialRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		MaterialRepository: NewPostgresMaterialRepository(db),
	}
}
