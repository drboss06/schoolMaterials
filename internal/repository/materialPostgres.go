package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"schoolMaterial/internal/models"
)

type PostgresMaterialRepository struct {
	DB *sqlx.DB
}

func NewPostgresMaterialRepository(db *sqlx.DB) *PostgresMaterialRepository {
	return &PostgresMaterialRepository{
		DB: db,
	}
}

func (r *PostgresMaterialRepository) CreateMaterial(m models.Material) error {

	query := fmt.Sprintf("INSERT INTO %s (uuid, type, status, title, content, created_at, updated_at) "+
		"values ($1, $2, $3, $4, $5, $6, $7)", "materials")

	_, err := r.DB.Exec(query, m.UUID, m.Type, m.Status, m.Title,
		m.Content, m.CreatedAt, m.UpdatedAt)

	return err
}

func (r *PostgresMaterialRepository) GetMaterialByUUID(uuid string) (*models.Material, error) {
	query := "SELECT * FROM materials WHERE uuid = $1"

	row := r.DB.QueryRow(query, uuid)
	material := &models.Material{}

	err := row.Scan(&material.UUID, &material.Type, &material.Status, &material.Title,
		&material.Content, &material.CreatedAt, &material.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return material, nil
}
