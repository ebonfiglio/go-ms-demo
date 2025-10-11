package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DisplayMenuOptions(options []string) string {
	fmt.Println("\nSelect an option:")
	for i, opt := range options {
		fmt.Printf("%d. %s\n", i+1, opt)
	}
	fmt.Print("Choice: ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func getEntityInput() string {
	fmt.Println("Enter entity as JSON:")
	reader := bufio.NewReader(os.Stdin)
	json, _ := reader.ReadString('\n')

	json = strings.TrimSpace(json)
	return json
}

func getId() int64 {
	fmt.Println("Enter Id:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid id:", err)
		return 0
	}

	return int64(id)
}
