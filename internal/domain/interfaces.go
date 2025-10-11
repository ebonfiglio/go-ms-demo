package domain

type UserRepository interface {
	InsertUser(u *User) (*User, error)
	GetAllUsers() ([]User, error)
	GetUser(id int64) (*User, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(id int64) (int64, error)
}

type OrganizationRepository interface {
	InsertOrganization(o *Organization) (*Organization, error)
	GetAll() ([]Organization, error)
	GetOrganization(id int64) (*Organization, error)
	UpdateOrganization(o *Organization) (*Organization, error)
	DeleteOrganization(id int64) (int64, error)
}

type JobRepository interface {
	InsertJob(j *Job) (*Job, error)
	GetAllJobs() ([]Job, error)
	GetJob(id int64) (*Job, error)
	UpdateJob(j *Job) (*Job, error)
	DeleteJob(id int64) (int64, error)
}

type UserService interface {
	CreateUser(u *User) (*User, error)
	GetAllUsers() ([]User, error)
	GetUser(id int64) (*User, error)
	UpdateUser(u *User) (*User, error)
	DeleteUser(id int64) (int64, error)
}

type OrganizationService interface {
	CreateOrganization(o *Organization) (*Organization, error)
	GetAllOrganizations() ([]Organization, error)
	GetOrganization(id int64) (*Organization, error)
	UpdateOrganization(o *Organization) (*Organization, error)
	DeleteOrganization(id int64) (int64, error)
}

type JobService interface {
	CreateJob(j *Job) (*Job, error)
	GetAllJobs() ([]Job, error)
	GetJob(id int64) (*Job, error)
	UpdateJob(j *Job) (*Job, error)
	DeleteJob(id int64) (int64, error)
}
