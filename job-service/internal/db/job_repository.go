package db

import (
	"fmt"
	"go-ms-demo/internal/domain"

	"github.com/jmoiron/sqlx"
)

type JobRepository struct {
	db *sqlx.DB
}

func NewJobRepository(db *sqlx.DB) *JobRepository {
	return &JobRepository{db}
}

func (r *JobRepository) InsertJob(j *domain.Job) (*domain.Job, error) {
	createdJob := &domain.Job{}
	err := r.db.Get(createdJob,
		`INSERT INTO jobs (name, organization_id)
     VALUES ($1, $2)
     RETURNING id, name, organization_id`,
		j.Name, j.OrganizationID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to insert job: %w", err)
	}
	return createdJob, nil
}

func (r *JobRepository) GetAllJobs() ([]domain.Job, error) {
	jobs := make([]domain.Job, 0)

	err := r.db.Select(&jobs, "SELECT id, name, organization_id from jobs")
	if err != nil {
		return nil, fmt.Errorf("failed to get all jobs: %w", err)
	}
	return jobs, nil
}

func (r *JobRepository) GetJob(id int64) (*domain.Job, error) {
	job := &domain.Job{}
	err := r.db.Get(job, "SELECT id, name, organization_id FROM jobs WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to get job: %w", err)
	}
	return job, nil
}

func (r *JobRepository) UpdateJob(j *domain.Job) (*domain.Job, error) {
	updatedJob := &domain.Job{}
	err := r.db.Get(updatedJob, "UPDATE jobs set name = $1, organization_id = $2 WHERE id = $3 RETURNING id, name, organization_id", j.Name, j.OrganizationID, j.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update job: %w", err)
	}
	return updatedJob, nil
}

func (r *JobRepository) DeleteJob(id int64) (int64, error) {
	result, err := r.db.Exec("DELETE FROM jobs WHERE id = $1", id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete job: %w", err)
	}
	return result.RowsAffected()
}
