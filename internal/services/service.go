package service

import (
	"schoolMaterial/internal/models"
	"schoolMaterial/internal/repository"
)

type MaterialServiceInterface interface {
	CreateMaterial(m models.Material) (uuid string, err error)
	GetMaterialByUUID(uuid string) (m *models.Material, err error)
	UpdateMaterial(uuid string, request models.UpdateRequest) (material *models.Material, err error)
	GetAllMaterials(active bool, materialType, startDate, endDate string) (materials []*models.Material, err error)
}

type Service struct {
	MaterialServiceInterface
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		MaterialServiceInterface: NewMaterialService(r.MaterialRepository),
	}
}
