package menu

import (
	"fmt"
	"go-db-demo/internal/domain"
)

func JobMenu(jobService domain.JobService) {
	for {
		choice := DisplayMenuOptions([]string{
			"Create Job",
			"Update Job",
			"Lookup Job",
			"Delete Job",
			"Get All Jobs",
			"Back",
		})

		switch choice {
		case "1":
			createJobCommand(jobService)
			fmt.Println("Success!")
		case "2":
			updateJobCommand(jobService)
			fmt.Println("Success!")
		case "3":
			getJobCommand(jobService)
		case "4":
			fmt.Println("Deleting Job...")
			deleteJobCommand(jobService)
		case "5":
			getAllJobsCommand(jobService)
		case "6":
			return
		}
	}
}

func createJobCommand(jobService domain.JobService) {
	newJobValues := getEntityInput()
	job, err := domain.JsonToJob(newJobValues)
	if err != nil {
		fmt.Println(err)
		return
	}
	job, err = jobService.CreateJob(job)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Job ID: ", job.ID)
}

func getAllJobsCommand(jobService domain.JobService) {
	jobs, err := jobService.GetAllJobs()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, j := range jobs {
		fmt.Println(j.ID, j.Name, j.OrganizationID)
	}
}

func getJobCommand(jobService domain.JobService) {
	id := getId()
	if id == 0 {
		return
	}
	job, err := jobService.GetJob(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(job.ID, job.Name, job.OrganizationID)

}

func updateJobCommand(jobService domain.JobService) {
	newJobValues := getEntityInput()
	job, err := domain.JsonToJob(newJobValues)
	if err != nil {
		fmt.Println(err)
		return
	}
	job, err = jobService.UpdateJob(job)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Job ID: ", job.ID)
	fmt.Println("Job Name: ", job.Name)
	fmt.Println("Job Org ID: ", job.OrganizationID)
}

func deleteJobCommand(jobService domain.JobService) {
	id := getId()
	if id == 0 {
		return
	}

	_, err := jobService.DeleteJob(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Job deleted!")
}
