package service

import (
	"fmt"
	"go-ms-demo/internal/domain"
)

type OrganizationService struct {
	orgRepo domain.OrganizationRepository
}

func NewOrganizationService(orgRepo domain.OrganizationRepository) *OrganizationService {
	return &OrganizationService{orgRepo: orgRepo}
}

func (s *OrganizationService) CreateOrganization(o *domain.Organization) (*domain.Organization, error) {
	org, err := s.orgRepo.InsertOrganization(o)
	if err != nil {
		return nil, fmt.Errorf("could not create organization: %w", err)
	}
	return org, nil
}

func (s *OrganizationService) GetAllOrganizations() ([]domain.Organization, error) {
	org, err := s.orgRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve all organizations: %w", err)
	}
	return org, nil
}

func (s *OrganizationService) GetOrganization(id int64) (*domain.Organization, error) {
	org, err := s.orgRepo.GetOrganization(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve organization: %w", err)
	}
	return org, nil
}

func (s *OrganizationService) UpdateOrganization(o *domain.Organization) (*domain.Organization, error) {
	org, err := s.orgRepo.UpdateOrganization(o)
	if err != nil {
		return nil, fmt.Errorf("could not update organization: %w", err)
	}
	return org, nil
}

func (s *OrganizationService) DeleteOrganization(id int64) (int64, error) {
	rowsAffected, err := s.orgRepo.DeleteOrganization(id)
	if err != nil {
		return 0, fmt.Errorf("could not delete organization: %w", err)
	}
	return rowsAffected, nil
}
