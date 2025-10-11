package domain

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	OrganizationID int64  `json:"organization_id" db:"organization_id"`
	JobID          int64  `json:"job_id" db:"job_id"`
}

func JsonToUser(input string) (*User, error) {
	var u User
	err := json.Unmarshal([]byte(input), &u)
	if err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}
	return &u, err
}
