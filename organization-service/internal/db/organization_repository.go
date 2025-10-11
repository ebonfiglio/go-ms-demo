package db

import (
	"fmt"
	"go-ms-demo/internal/domain"

	"github.com/jmoiron/sqlx"
)

type OrganizationRepository struct {
	db *sqlx.DB
}

func NewOrganizationRepository(db *sqlx.DB) *OrganizationRepository {
	return &OrganizationRepository{db}
}

func (r *OrganizationRepository) InsertOrganization(o *domain.Organization) (*domain.Organization, error) {
	createdOrganization := &domain.Organization{}
	err := r.db.Get(createdOrganization,
		"INSERT INTO organizations (name) VALUES ($1) RETURNING id, name",
		o.Name,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert organization: %w", err)
	}
	return createdOrganization, nil
}

func (r *OrganizationRepository) GetAll() ([]domain.Organization, error) {
	organizations := make([]domain.Organization, 0)
	err := r.db.Select(&organizations, "select id, name from organizations")
	if err != nil {
		return nil, fmt.Errorf("failed to get all organizations: %w", err)
	}
	return organizations, nil
}

func (r *OrganizationRepository) GetOrganization(id int64) (*domain.Organization, error) {
	organization := &domain.Organization{}
	err := r.db.Get(organization, "SELECT id, name FROM organizations WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization: %w", err)
	}
	return organization, nil
}

func (r *OrganizationRepository) UpdateOrganization(o *domain.Organization) (*domain.Organization, error) {
	updatedOrganization := &domain.Organization{}

	err := r.db.Get(updatedOrganization, "UPDATE organizations SET name = $1 WHERE id = $2 RETURNING id, name", o.Name, o.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update organization: %w", err)
	}

	return updatedOrganization, nil

}

func (r OrganizationRepository) DeleteOrganization(id int64) (int64, error) {
	result, err := r.db.Exec("DELETE FROM organizations WHERE id = $1", id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete organization: %w", err)
	}

	return result.RowsAffected()
}
