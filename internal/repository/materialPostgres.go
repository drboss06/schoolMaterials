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

// CreateMaterial inserts a new material into the "materials" table in the Postgres database.
//
// Parameters:
// - m: a models.Material struct representing the material to be created.
//
// Returns:
// - error: an error if the insertion fails, otherwise nil.
func (r *PostgresMaterialRepository) CreateMaterial(m models.Material) error {

	query := fmt.Sprintf("INSERT INTO %s (uuid, type, status, title, content, created_at, updated_at) "+
		"values ($1, $2, $3, $4, $5, $6, $7)", "materials")

	_, err := r.DB.Exec(query, m.UUID, m.Type, m.Status, m.Title,
		m.Content, m.CreatedAt, m.UpdatedAt)

	return err
}

// GetMaterialByUUID retrieves a material from the "materials" table in the Postgres database based on the provided UUID.
//
// Parameters:
// - uuid: a string representing the UUID of the material to retrieve.
//
// Returns:
// - *models.Material: a pointer to a models.Material struct representing the retrieved material, or nil if no material is found.
// - error: an error if the retrieval fails, or nil if the retrieval is successful.
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

// UpdateMaterial updates a material in the "materials" table in the Postgres database.
//
// Parameters:
// - material: a models.Material struct representing the material to be updated.
//
// Returns:
// - error: an error if the update fails, otherwise nil.
func (r *PostgresMaterialRepository) UpdateMaterial(material models.Material) error {
	query := fmt.Sprintf("UPDATE %s SET status = $1, title = $2, content = $3, updated_at = $4 WHERE uuid = $5", "materials")

	_, err := r.DB.Exec(query, material.Status, material.Title, material.Content, material.UpdatedAt, material.UUID)

	return err
}

// GetAllMaterials retrieves all materials from the "materials" table in the Postgres database.
//
// Parameters:
// - active: a boolean indicating whether to retrieve only active materials.
// - materialType: a string representing the type of materials to retrieve.
// - startDate: a string representing the start date of materials to retrieve.
// - endDate: a string representing the end date of materials to retrieve.
//
// Returns:
// - []*models.Material: a slice of pointers to models.Material structs representing the retrieved materials.
// - error: an error if the retrieval fails, otherwise nil.
func (r *PostgresMaterialRepository) GetAllMaterials(active bool, materialType, startDate, endDate string) ([]*models.Material, error) {
	query := "SELECT uuid, type, title, created_at, updated_at FROM materials WHERE 1=1"

	var args []interface{}
	if active {
		query += " AND status = 'active'"
	}

	argIndex := 1

	if materialType != "" {
		query += fmt.Sprintf(" AND type = $%d", argIndex)
		args = append(args, materialType)
		argIndex++
	}

	if startDate != "" && endDate != "" {
		query += fmt.Sprintf(" AND created_at BETWEEN $%d AND $%d", argIndex, argIndex+1)
		args = append(args, startDate, endDate)
		argIndex += 2
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var materials []*models.Material
	for rows.Next() {
		material := &models.Material{}
		err := rows.Scan(&material.UUID, &material.Type, &material.Title, &material.CreatedAt, &material.UpdatedAt)
		if err != nil {
			return nil, err
		}
		materials = append(materials, material)
	}
	return materials, nil
}
