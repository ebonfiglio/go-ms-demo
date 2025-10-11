package menu

import (
	"fmt"
	"go-db-demo/internal/domain"
)

func OrganizationMenu(orgService domain.OrganizationService) {
	for {
		choice := DisplayMenuOptions([]string{
			"Create Org",
			"Update Org",
			"Lookup Org",
			"Delete Org",
			"List Orgs",
			"Back",
		})

		switch choice {
		case "1":
			createOrganizationCommand(orgService)
			fmt.Println("Success!")
		case "2":
			updateOrganizationCommand(orgService)
			fmt.Println("Success!")
		case "3":
			getOrganizationCommand(orgService)
		case "4":
			fmt.Println("Deleting Org...")
			deleteOrganizationCommand(orgService)
		case "5":
			listOrganizationsCommand(orgService)
		case "6":
			return
		}
	}
}

func createOrganizationCommand(orgService domain.OrganizationService) {
	newOrgValues := getEntityInput()
	org, err := domain.JsonToOrganization(newOrgValues)
	if err != nil {
		fmt.Println(err)
		return
	}
	org, err = orgService.CreateOrganization(org)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Organziation ID: ", org.ID)
}

func listOrganizationsCommand(orgService domain.OrganizationService) {
	organizations, err := orgService.GetAllOrganizations()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, o := range organizations {
		fmt.Println(o.ID, o.Name)
	}
}

func getOrganizationCommand(orgService domain.OrganizationService) {
	id := getId()
	if id == 0 {
		return
	}
	organization, err := orgService.GetOrganization(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(organization.ID, organization.Name)
}

func updateOrganizationCommand(orgService domain.OrganizationService) {
	newOrgValues := getEntityInput()
	org, err := domain.JsonToOrganization(newOrgValues)
	if err != nil {
		fmt.Println(err)
		return
	}
	org, err = orgService.UpdateOrganization(org)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Organziation ID: ", org.ID)
	fmt.Println("Organziation Name: ", org.Name)
}

func deleteOrganizationCommand(orgService domain.OrganizationService) {
	id := getId()
	if id == 0 {
		return
	}

	_, err := orgService.DeleteOrganization(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Organization deleted!")
}
