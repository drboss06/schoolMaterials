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

// CreateMaterial creates a new material in the MaterialService.
//
// Parameters:
// - m: the material to be created.
//
// Returns:
// - string: the UUID of the created material.
// - error: an error if the material type or status is not allowed, or if there was an error creating the material.
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

// GetMaterialByUUID retrieves a material from the MaterialService by its UUID.
//
// Parameters:
// - uuid: the UUID of the material to retrieve.
//
// Returns:
// - *models.Material: the material with the given UUID, or nil if not found.
// - error: an error if there was a problem retrieving the material.
func (s *MaterialService) GetMaterialByUUID(uuid string) (*models.Material, error) {
	mayerial, err := s.repo.GetMaterialByUUID(uuid)

	if err != nil {
		return nil, err
	}

	return mayerial, nil
}

// UpdateMaterial updates a material with the given UUID in the MaterialService.
//
// Parameters:
// - uuid: the UUID of the material to update.
// - request: the UpdateRequest containing the new status, title, and content for the material.
//
// Returns:
// - *models.Material: the updated material, or nil if there was an error.
// - error: an error if there was a problem updating the material.
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

// GetAllMaterials retrieves all materials from the MaterialService based on the provided criteria.
//
// Parameters:
// - active: a boolean indicating whether to retrieve only active materials.
// - materialType: a string representing the type of materials to retrieve.
// - startDate: a string representing the start date of the materials to retrieve.
// - endDate: a string representing the end date of the materials to retrieve.
//
// Returns:
// - materials: a slice of pointers to Material objects representing the retrieved materials.
// - err: an error if there was a problem retrieving the materials.
func (s *MaterialService) GetAllMaterials(active bool, materialType, startDate, endDate string) (materials []*models.Material, err error) {
	materials, err = s.repo.GetAllMaterials(active, materialType, startDate, endDate)

	if err != nil {
		return nil, err
	}

	return materials, nil
}
