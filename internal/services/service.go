package service

import (
	"schoolMaterial/internal/models"
	"schoolMaterial/internal/repository"
)

type MaterialServiceInterface interface {
	CreateMaterial(m models.Material) (uuid string, err error)
	GetMaterialByUUID(uuid string) (m *models.Material, err error)
}

type Service struct {
	MaterialServiceInterface
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		MaterialServiceInterface: NewMaterialService(r.MaterialRepository),
	}
}
