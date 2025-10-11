package main

import (
	"fmt"
	"go-db-demo/console/menu"
	"go-db-demo/internal/db"
	"go-db-demo/internal/service"
)

func main() {

	dbConn := db.Connect()
	defer dbConn.Close()

	userRepo := db.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)

	jobRepo := db.NewJobRepository(dbConn)
	jobService := service.NewJobService(jobRepo)

	orgRepo := db.NewOrganizationRepository(dbConn)
	orgService := service.NewOrganizationService(orgRepo)

	fmt.Println("Welcome to the Management System")

	for {
		choice := menu.DisplayMenuOptions([]string{"Organizations", "Jobs", "Users", "Exit"})

		switch choice {
		case "1":
			menu.OrganizationMenu(orgService)
		case "2":
			menu.JobMenu(jobService)
		case "3":
			menu.UserMenu(userService)
		case "4":
			fmt.Println("Goodbye!")
			return
		}
	}
}
