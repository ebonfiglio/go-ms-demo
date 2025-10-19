package db

import (
	"fmt"
	"go-ms-demo/user-service/internal/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) InsertUser(u *domain.User) (*domain.User, error) {
	createdUser := &domain.User{}
	err := r.db.Get(createdUser,
		"INSERT INTO users (name, job_id, organization_id) VALUES ($1, $2, $3) RETURNING id, name, job_id, organization_id",
		u.Name, u.JobID, u.OrganizationID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}
	return createdUser, nil
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	users := make([]domain.User, 0)
	err := r.db.Select(&users, "SELECT id, name, job_id, organization_id from users")
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	return users, nil
}

func (r *UserRepository) GetUser(id int64) (*domain.User, error) {
	user := &domain.User{}
	err := r.db.Get(user, "SELECT id, name, job_id, organization_id FROM users WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(u *domain.User) (*domain.User, error) {
	updatedUser := &domain.User{}
	err := r.db.Get(updatedUser, "UPDATE users set name = $1, job_id = $2, organization_id = $3 WHERE id = $4 RETURNING id, name, job_id, organization_id", u.Name, u.JobID, u.OrganizationID, u.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}
	return updatedUser, nil
}

func (r *UserRepository) DeleteUser(id int64) (int64, error) {
	result, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete user: %w", err)
	}
	return result.RowsAffected()
}
