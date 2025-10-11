package menu

import (
	"fmt"
	"go-db-demo/internal/domain"
)

func UserMenu(userService domain.UserService) {
	for {
		choice := DisplayMenuOptions([]string{
			"Create User",
			"Update User",
			"Lookup User",
			"Delete User",
			"Get all Users",
			"Back",
		})

		switch choice {
		case "1":
			createUserCommand(userService)
			fmt.Println("Success!")
		case "2":
			updateUserCommand(userService)
			fmt.Println("Success!")
		case "3":
			getUserCommand(userService)
		case "4":
			fmt.Println("Deleting User...")
			deleteUserCommand(userService)
		case "5":
			getAllUsersCommand(userService)
		case "6":
			return
		}
	}
}

func createUserCommand(userService domain.UserService) {
	newUserValues := getEntityInput()
	user, err := domain.JsonToUser(newUserValues)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err = userService.CreateUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User ID: ", user.ID)
}

func getAllUsersCommand(userService domain.UserService) {
	users, err := userService.GetAllUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.JobID, u.OrganizationID)
	}
}

func getUserCommand(userService domain.UserService) {
	id := getId()
	if id == 0 {
		return
	}
	user, err := userService.GetUser(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user.ID, user.Name, user.JobID, user.OrganizationID)

}

func updateUserCommand(userService domain.UserService) {
	newUserValues := getEntityInput()
	user, err := domain.JsonToUser(newUserValues)
	if err != nil {
		fmt.Println(err)
		return
	}
	user, err = userService.UpdateUser(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User ID: ", user.ID)
	fmt.Println("User Name: ", user.Name)
	fmt.Println("User Job ID: ", user.JobID)
	fmt.Println("User Org ID: ", user.OrganizationID)
}

func deleteUserCommand(userService domain.UserService) {
	id := getId()
	if id == 0 {
		return
	}

	_, err := userService.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User deleted!")
}
