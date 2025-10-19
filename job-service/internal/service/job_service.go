package service

import (
	"fmt"
	"go-ms-demo/job-service/internal/domain"
)

type JobService struct {
	jobRepo domain.JobRepository
}

func NewJobService(jobRepo domain.JobRepository) *JobService {
	return &JobService{jobRepo: jobRepo}
}

func (s *JobService) CreateJob(j *domain.Job) (*domain.Job, error) {
	job, err := s.jobRepo.InsertJob(j)
	if err != nil {
		return nil, fmt.Errorf("could not create job: %w", err)
	}
	return job, nil
}

func (s *JobService) GetAllJobs() ([]domain.Job, error) {
	jobs, err := s.jobRepo.GetAllJobs()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve all jobs: %w", err)
	}

	return jobs, nil
}

func (s *JobService) GetJob(id int64) (*domain.Job, error) {
	job, err := s.jobRepo.GetJob(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve job: %w", err)
	}
	return job, nil
}

func (s *JobService) UpdateJob(j *domain.Job) (*domain.Job, error) {
	job, err := s.jobRepo.UpdateJob(j)
	if err != nil {
		return nil, fmt.Errorf("could not update job: %w", err)
	}
	return job, nil
}

func (s *JobService) DeleteJob(id int64) (int64, error) {
	rowsAffected, err := s.jobRepo.DeleteJob(id)
	if err != nil {
		return 0, fmt.Errorf("could not delete job: %w", err)
	}
	return rowsAffected, nil
}
