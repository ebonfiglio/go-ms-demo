package domain

type JobRepository interface {
	InsertJob(j *Job) (*Job, error)
	GetAllJobs() ([]Job, error)
	GetJob(id int64) (*Job, error)
	UpdateJob(j *Job) (*Job, error)
	DeleteJob(id int64) (int64, error)
}

type JobService interface {
	CreateJob(j *Job) (*Job, error)
	GetAllJobs() ([]Job, error)
	GetJob(id int64) (*Job, error)
	UpdateJob(j *Job) (*Job, error)
	DeleteJob(id int64) (int64, error)
}
