package service

import (
	"errors"
	"github.com/google/uuid"
	"schoolMaterial/internal/models"
	"schoolMaterial/internal/repository"
	"time"
)

type MaterialService struct {
	repo repository.MaterialRepository
}

var allowedMaterialTypes = map[string]bool{
	"статья":      true,
	"видеоролик":  true,
	"презентация": true,
}

var allowedMaterialStatuses = map[string]bool{
	"архивный": true,
	"активный": true,
}

func NewMaterialService(r repository.MaterialRepository) *MaterialService {
	return &MaterialService{
		repo: r,
	}
}

func (s *MaterialService) CreateMaterial(m models.Material) (string, error) {

	if !allowedMaterialTypes[m.Type] {
		return "", errors.New("Not allowed material type")
	}

	if !allowedMaterialStatuses[m.Status] {
		return "", errors.New("Not allowed material status")
	}

	m.UUID = uuid.New().String()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()

	if err := s.repo.CreateMaterial(m); err != nil {
		return "", err
	}

	return m.UUID, nil
}

func (s *MaterialService) GetMaterialByUUID(uuid string) (*models.Material, error) {
	mayerial, err := s.repo.GetMaterialByUUID(uuid)

	if err != nil {
		return nil, err
	}

	return mayerial, nil
}

func (s *MaterialService) UpdateMaterial(uuid string, request models.UpdateRequest) (material *models.Material, err error) {
	material, err = s.repo.GetMaterialByUUID(uuid)

	if err != nil {
		return nil, err
	}

	material.Status = request.Status
	material.Title = request.Title
	material.Content = request.Content
	material.UpdatedAt = time.Now()

	if err := s.repo.UpdateMaterial(*material); err != nil {
		return nil, err
	}

	return material, nil
}

func (s *MaterialService) GetAllMaterials(active bool, materialType, startDate, endDate string) (materials []*models.Material, err error) {
	materials, err = s.repo.GetAllMaterials(active, materialType, startDate, endDate)

	if err != nil {
		return nil, err
	}

	return materials, nil
}
