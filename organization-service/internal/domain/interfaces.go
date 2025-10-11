package domain

type OrganizationRepository interface {
	InsertOrganization(o *Organization) (*Organization, error)
	GetAll() ([]Organization, error)
	GetOrganization(id int64) (*Organization, error)
	UpdateOrganization(o *Organization) (*Organization, error)
	DeleteOrganization(id int64) (int64, error)
}

type OrganizationService interface {
	CreateOrganization(o *Organization) (*Organization, error)
	GetAllOrganizations() ([]Organization, error)
	GetOrganization(id int64) (*Organization, error)
	UpdateOrganization(o *Organization) (*Organization, error)
	DeleteOrganization(id int64) (int64, error)
}
